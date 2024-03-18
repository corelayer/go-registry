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

package certificates

import (
	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type Passphrase struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Value string `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty" secure:"true"`
}

func (c Passphrase) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Passphrase{},
		Encrypted: SecurePassphrase{},
	}
}

type SecurePassphrase struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Value        string                    `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecurePassphrase) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Passphrase{},
		Encrypted: SecurePassphrase{},
	}
}

func (s SecurePassphrase) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}