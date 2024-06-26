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

package registry

import (
	"fmt"

	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
	"github.com/corelayer/go-netscaleradc-nitro/pkg/nitro"
	"golang.org/x/crypto/ssh"
)

type NetScalerAdcEnvironment struct {
	Name        string                   `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`                     // Target environment name, such as "Production"
	Management  NetScalerAdcNode         `json:"management,omitempty" yaml:"management,omitempty" mapstructure:"management,omitempty" secure:"true"`    // Connection details for the Management Address (SNIP / Cluster IP) of the environment
	Nodes       []NetScalerAdcNode       `json:"nodes,omitempty" yaml:"nodes,omitempty" mapstructure:"nodes,omitempty" secure:"true"`                   // Connection details for the individual Nodes of each node
	Credentials []NetScalerAdcCredential `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"` // Connection credentials
	Settings    NetScalerAdcSettings     `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings,omitempty" secure:"false"`         // Connection settings for Nitro Client
}

func (e NetScalerAdcEnvironment) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcEnvironment{},
		Encrypted: SecureNetScalerAdcEnvironment{},
	}
}

func (e NetScalerAdcEnvironment) GetCredentialByName(name string) (NetScalerAdcCredential, error) {
	for _, c := range e.Credentials {
		if c.Name == name {
			return c, nil
		}
	}
	return NetScalerAdcCredential{}, fmt.Errorf("could not find credential %s in environment %s", name, e.Name)
}

func (e NetScalerAdcEnvironment) GetNodeScpClient(nodeName string, credential NetScalerAdcCredential, f ssh.HostKeyCallback) (scp.Client, error) {
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

func (e NetScalerAdcEnvironment) GetNodeNitroClient(nodeName string, credential NetScalerAdcCredential) (*nitro.Client, error) {
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

func (e NetScalerAdcEnvironment) GetManagementClient(credential NetScalerAdcCredential) (*nitro.Client, error) {
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

func (e NetScalerAdcEnvironment) GetNodeNames() []string {
	var (
		output []string
	)
	output = make([]string, len(e.Nodes))
	for i := 0; i < len(e.Nodes); i++ {
		output[i] = e.Nodes[i].Name
	}

	return output
}

func (e NetScalerAdcEnvironment) GetNodes() []NetScalerAdcNode {
	var output []NetScalerAdcNode
	output = append(output, e.Nodes...)

	if e.HasManagement() {
		output = append(output, e.Management)
	}
	return output
}

func (e NetScalerAdcEnvironment) GetPrimaryClient(credential NetScalerAdcCredential) (*nitro.Client, error) {
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

func (e NetScalerAdcEnvironment) HasNodes() bool {
	if len(e.Nodes) == 0 {
		return false
	}
	return true
}

func (e NetScalerAdcEnvironment) HasManagement() bool {
	emptyNode := NetScalerAdcNode{}
	if e.Management != emptyNode {
		return true
	}
	return false
}

type SecureNetScalerAdcEnvironment struct {
	Name         string                         `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`                     // Target environment name, such as "Production"
	Management   SecureNetScalerAdcNode         `json:"management,omitempty" yaml:"management,omitempty" mapstructure:"management,omitempty" secure:"true"`    // Connection details for the Management Address (SNIP / Cluster IP) of the environment
	Nodes        []SecureNetScalerAdcNode       `json:"nodes,omitempty" yaml:"nodes,omitempty" mapstructure:"nodes,omitempty" secure:"true"`                   // Connection details for the individual Nodes of each node
	Credentials  []SecureNetScalerAdcCredential `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"` // Connection credentials
	Settings     NetScalerAdcSettings           `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings,omitempty" secure:"false"`         // Connection settings for Nitro Client
	CryptoParams cryptostruct.CryptoParams      `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (e SecureNetScalerAdcEnvironment) GetCredentialByName(name string) (SecureNetScalerAdcCredential, error) {
	for _, c := range e.Credentials {
		if c.Name == name {
			return c, nil
		}
	}
	return SecureNetScalerAdcCredential{}, fmt.Errorf("could not find credential %s in environment %s", name, e.Name)
}

func (s SecureNetScalerAdcEnvironment) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureNetScalerAdcEnvironment) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcEnvironment{},
		Encrypted: SecureNetScalerAdcEnvironment{},
	}
}
