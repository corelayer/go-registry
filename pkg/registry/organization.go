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
	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type Organization struct {
	Name     string               `json:"name" yaml:"name" mapstructure:"name" secure:"false"`
	Registry OrganizationRegistry `json:"registry,omitempty" yaml:"registry,omitempty" mapstructure:"registry,omitempty" secure:"true"`
}

func (o Organization) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Organization{},
		Encrypted: SecureOrganization{},
	}
}

type SecureOrganization struct {
	Name         string                     `json:"name" yaml:"name" mapstructure:"name" secure:"false"`
	Registry     SecureOrganizationRegistry `json:"registry,omitempty" yaml:"registry,omitempty" mapstructure:"registry,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams  `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureOrganization) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureOrganization) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Organization{},
		Encrypted: SecureOrganization{},
	}
}
