/*
 * Copyright 2024 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package adc

import (
	"fmt"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"golang.org/x/crypto/ssh"
)

type Environment struct {
	Name        string       `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`                     // Target environment name, such as "Production"
	Management  Node         `json:"management,omitempty" yaml:"management,omitempty" mapstructure:"management,omitempty" secure:"true"`    // Connection details for the Management Address (SNIP / Cluster IP) of the environment
	Nodes       []Node       `json:"nodes,omitempty" yaml:"nodes,omitempty" mapstructure:"nodes,omitempty" secure:"true"`                   // Connection details for the individual Nodes of each node
	Credentials []Credential `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"` // Connection credentials
	Settings    Settings     `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings,omitempty" secure:"false"`         // Connection settings for Nitro Client
}

func (e Environment) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Environment{},
		Encrypted: SecureEnvironment{},
	}
}

func (e Environment) GetCredentialByName(name string) (Credential, error) {
	for _, c := range e.Credentials {
		if c.Name == name {
			return c, nil
		}
	}
	return Credential{}, fmt.Errorf("could not find credential %s in environment %s", name, e.Name)
}

func (e Environment) GetNodeScpClient(nodeName string, credential Credential, f ssh.HostKeyCallback) (scp.Client, error) {
	var (
		err          error
		clientConfig ssh.ClientConfig
	)

	for _, n := range e.GetNodes() {
		if n.Name == nodeName {
			clientConfig, err = auth.PasswordKey(credential.Username, credential.Password, f)
			if err != nil {
				return scp.Client{}, err
			}

			return scp.NewClient(n.Address+":22", &clientConfig), nil
		}
	}
	return scp.Client{}, fmt.Errorf("could not intialize scp client")
}

func (e Environment) GetNodeNitroClient(nodeName string, credential Credential) (*nitro.Client, error) {
	var (
		err    error
		client *nitro.Client
	)

	nitroCredential := nitro.Credentials{
		Username: credential.Username,
		Password: credential.Password,
	}

	nitroSettings := nitro.ConnectionSettings{
		UseSsl:                    e.Settings.UseSsl,
		Timeout:                   e.Settings.Timeout,
		UserAgent:                 e.Settings.UserAgent,
		ValidateServerCertificate: e.Settings.ValidateServerCertificate,
		LogTlsSecrets:             e.Settings.LogTlsSecrets,
		LogTlsSecretsDestination:  e.Settings.LogTlsSecretsDestination,
		AutoLogin:                 e.Settings.AutoLogin,
	}

	for _, n := range e.GetNodes() {
		if n.Name == nodeName {
			client, err = nitro.NewClient(n.Name, n.Address, nitroCredential, nitroSettings)
			if err != nil {
				return nil, fmt.Errorf("could not create client for node %s with error %w", nodeName, err)
			}
			return client, nil
		}
	}
	return nil, fmt.Errorf("could not create client for node %s with error: node not found in environment %s", nodeName, e.Name)
}

func (e Environment) GetManagementClient(credential Credential) (*nitro.Client, error) {
	// Return the SNIP Node if defined in the environment
	if !e.HasManagement() {
		return nil, fmt.Errorf("no management node defined for environment %s", e.Name)

	}

	nitroCredential := nitro.Credentials{
		Username: credential.Username,
		Password: credential.Password,
	}

	nitroSettings := nitro.ConnectionSettings{
		UseSsl:                    e.Settings.UseSsl,
		Timeout:                   e.Settings.Timeout,
		UserAgent:                 e.Settings.UserAgent,
		ValidateServerCertificate: e.Settings.ValidateServerCertificate,
		LogTlsSecrets:             e.Settings.LogTlsSecrets,
		LogTlsSecretsDestination:  e.Settings.LogTlsSecretsDestination,
		AutoLogin:                 e.Settings.AutoLogin,
	}

	client, err := nitro.NewClient(e.Management.Name, e.Management.Address, nitroCredential, nitroSettings)
	if err != nil {
		return nil, fmt.Errorf("could not create client for management node %s for environment %s with error %w", e.Management.Name, e.Name, err)
	}
	return client, nil
}

func (e Environment) GetNodeNames() []string {
	var (
		output []string
	)
	output = make([]string, len(e.Nodes))
	for i := 0; i < len(e.Nodes); i++ {
		output[i] = e.Nodes[i].Name
	}

	return output
}

func (e Environment) GetNodes() []Node {
	var output []Node
	output = append(output, e.Nodes...)

	if e.HasManagement() {
		output = append(output, e.Management)
	}
	return output
}

func (e Environment) GetPrimaryClient(credential Credential) (*nitro.Client, error) {
	var (
		err    error
		client *nitro.Client
	)

	client, _ = e.GetManagementClient(credential)
	if client != nil {
		return client, nil
	}

	if !e.HasNodes() {
		return nil, fmt.Errorf("no individual nodes defined for environment %s", e.Name)
	}

	// Loop over individual nodes to determine which one is primary
	// It is not guaranteed in a High-Available environment that one of the nodes is automatically primary, so we have to iterate over both
	for _, n := range e.Nodes {
		client, err = e.GetNodeNitroClient(n.Name, credential)
		if err != nil {
			return nil, fmt.Errorf("could not create client for node %s to determine status as primary node for environment %s with error %w", n.Name, e.Name, err)
		}

		var isPrimary bool
		if isPrimary, err = client.IsPrimaryNode(); err != nil {
			return nil, fmt.Errorf("could not determine status for node %s for environment %s with error %w", n.Name, e.Name, err)
		}
		if isPrimary {
			return client, nil
		}
	}

	return nil, fmt.Errorf("could not find a primary node for environment %s", e.Name)
}

func (e Environment) HasNodes() bool {
	if len(e.Nodes) == 0 {
		return false
	}
	return true
}

func (e Environment) HasManagement() bool {
	emptyNode := Node{}
	if e.Management != emptyNode {
		return true
	}
	return false
}

type SecureEnvironment struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`                     // Target environment name, such as "Production"
	Management   SecureNode                `json:"management,omitempty" yaml:"management,omitempty" mapstructure:"management,omitempty" secure:"true"`    // Connection details for the Management Address (SNIP / Cluster IP) of the environment
	Nodes        []SecureNode              `json:"nodes,omitempty" yaml:"nodes,omitempty" mapstructure:"nodes,omitempty" secure:"true"`                   // Connection details for the individual Nodes of each node
	Credentials  []SecureCredential        `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"` // Connection credentials
	Settings     Settings                  `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings,omitempty" secure:"false"`         // Connection settings for Nitro Client
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureEnvironment) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Environment{},
		Encrypted: SecureEnvironment{},
	}
}

func (s SecureEnvironment) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
