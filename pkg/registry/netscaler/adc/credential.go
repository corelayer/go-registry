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

import "github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

type Credential struct {
	Name     string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Username string `json:"username,omitempty" yaml:"username,omitempty" mapstructure:"username,omitempty" secure:"true"`
	Password string `json:"password,omitempty" yaml:"password,omitempty" mapstructure:"password,omitempty" secure:"true"`
}

func (c Credential) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Credential{},
		Encrypted: SecureCredential{},
	}
}

type SecureCredential struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Username     string                    `json:"username,omitempty" yaml:"username,omitempty" mapstructure:"username,omitempty" secure:"true"`
	Password     string                    `json:"password,omitempty" yaml:"password,omitempty" mapstructure:"password,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureCredential) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Credential{},
		Encrypted: SecureCredential{},
	}
}

func (s SecureCredential) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
