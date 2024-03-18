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

package netscaler

import (
	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

	"github.com/corelayer/go-registry/pkg/registry/machines/netscaler/adc"
	"github.com/corelayer/go-registry/pkg/registry/machines/netscaler/sdx"
)

type Registry struct {
	Adc adc.Registry `json:"adc,omitempty" yaml:"adc,omitempty" mapstructure:"adc,omitempty" secure:"true"`
	Sdx sdx.Registry `json:"sdx,omitempty" yaml:"sdx,omitempty" mapstructure:"sdx,omitempty" secure:"true"`
}

func (r Registry) GetAdcEnvironmentByName(name string) (adc.Environment, error) {
	return r.Adc.GetEnvironmentByName(name)
}

func (r Registry) GetAdcEnvironmentNames() []string {
	return r.Adc.GetEnvironmentNames()
}

func (r Registry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

type SecureRegistry struct {
	Adc          adc.SecureRegistry        `json:"adc,omitempty" yaml:"adc,omitempty" mapstructure:"adc,omitempty" secure:"true"`
	Sdx          sdx.SecureRegistry        `json:"sdx,omitempty" yaml:"sdx,omitempty" mapstructure:"sdx,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

func (s SecureRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
