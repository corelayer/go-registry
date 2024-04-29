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

func NewNetScalerRegistry() NetScalerRegistry {
	return NetScalerRegistry{
		Adc: NewNetScalerAdcRegistry(),
		Sdx: NewNetScalerSdxRegistry(),
	}
}

type NetScalerRegistry struct {
	Adc NetScalerAdcRegistry `json:"adc,omitempty" yaml:"adc,omitempty" mapstructure:"adc,omitempty" secure:"true"`
	Sdx NetScalerSdxRegistry `json:"sdx,omitempty" yaml:"sdx,omitempty" mapstructure:"sdx,omitempty" secure:"true"`
}

func (r NetScalerRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerRegistry{},
		Encrypted: SecureNetScalerRegistry{},
	}
}

type SecureNetScalerRegistry struct {
	Adc          SecureNetScalerAdcRegistry `json:"adc,omitempty" yaml:"adc,omitempty" mapstructure:"adc,omitempty" secure:"true"`
	Sdx          SecureNetScalerSdxRegistry `json:"sdx,omitempty" yaml:"sdx,omitempty" mapstructure:"sdx,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams  `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureNetScalerRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureNetScalerRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerRegistry{},
		Encrypted: SecureNetScalerRegistry{},
	}
}
