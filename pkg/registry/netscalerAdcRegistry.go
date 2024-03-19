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

type NetScalerAdcRegistry struct {
	Environments []NetScalerAdcEnvironment `json:"environments,omitempty" yaml:"environments,omitempty" mapstructure:"environments,omitempty" secure:"true"`
}

func (r NetScalerAdcRegistry) GetEnvironmentByName(name string) (NetScalerAdcEnvironment, error) {
	for _, e := range r.Environments {
		if e.Name == name {
			return e, nil
		}
	}
	return NetScalerAdcEnvironment{}, NewItemNotFoundError("netscaler adc environment", name)
}

func (r NetScalerAdcRegistry) GetEnvironmentNames() []string {
	names := make([]string, len(r.Environments))
	for i := 0; i < len(r.Environments); i++ {
		names[i] = r.Environments[i].Name
	}
	return names
}

func (r NetScalerAdcRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcRegistry{},
		Encrypted: SecureNetScalerAdcRegistry{},
	}
}

type SecureNetScalerAdcRegistry struct {
	Environments []SecureNetScalerAdcEnvironment `json:"environments,omitempty" yaml:"environments,omitempty" mapstructure:"environments,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams       `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureNetScalerAdcRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (r SecureNetScalerAdcRegistry) GetEnvironmentByName(name string) (SecureNetScalerAdcEnvironment, error) {
	for _, e := range r.Environments {
		if e.Name == name {
			return e, nil
		}
	}
	return SecureNetScalerAdcEnvironment{}, NewItemNotFoundError("netscaler adc environment", name)
}

func (r SecureNetScalerAdcRegistry) GetEnvironmentNames() []string {
	names := make([]string, len(r.Environments))
	for i := 0; i < len(r.Environments); i++ {
		names[i] = r.Environments[i].Name
	}
	return names
}

func (s SecureNetScalerAdcRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcRegistry{},
		Encrypted: SecureNetScalerAdcRegistry{},
	}
}
