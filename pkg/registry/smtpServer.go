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

import "github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

type SmtpServerSettings struct {
	UseSsl     bool `json:"useSsl,omitempty" yaml:"useSsl,omitempty" mapstructure:"useSsl,omitempty" secure:"false"`
	UseTls     bool `json:"useTls,omitempty" yaml:"useTls,omitempty" mapstructure:"useTls,omitempty" secure:"false"`
	CustomPort int  `json:"customPort,omitempty" yaml:"customPort,omitempty" mapstructure:"customPort,omitempty" secure:"false"`
}

type SmtpServer struct {
	Name        string             `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address     string             `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
	Credentials SmtpCredentials    `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"`
	Settings    SmtpServerSettings `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings" secure:"false"`
}

func (s SmtpServer) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: SmtpServer{},
		Encrypted: SecureSmtpServer{},
	}
}

type SecureSmtpServer struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address      string                    `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
	Credentials  SecureSmtpCredentials     `json:"credentials,omitempty" yaml:"credentials,omitempty" mapstructure:"credentials,omitempty" secure:"true"`
	Settings     SmtpServerSettings        `json:"settings,omitempty" yaml:"settings,omitempty" mapstructure:"settings" secure:"false"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureSmtpServer) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureSmtpServer) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: SmtpServer{},
		Encrypted: SecureSmtpServer{},
	}
}
