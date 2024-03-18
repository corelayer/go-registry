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

package acme

import "github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

type User struct {
	Name                   string                 `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Email                  string                 `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty" secure:"true"`
	ExternalAccountBinding ExternalAccountBinding `json:"eab,omitempty" yaml:"eab,omitempty" mapstructure:"eab,omitempty" secure:"true"`
}

func (u User) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: User{},
		Encrypted: SecureUser{},
	}
}

type SecureUser struct {
	Name                   string                       `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Email                  string                       `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty" secure:"true"`
	ExternalAccountBinding SecureExternalAccountBinding `json:"eab,omitempty" yaml:"eab,omitempty" mapstructure:"eab,omitempty" secure:"true"`
	CryptoParams           cryptostruct.CryptoParams    `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureUser) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: User{},
		Encrypted: SecureUser{},
	}
}

func (s SecureUser) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
