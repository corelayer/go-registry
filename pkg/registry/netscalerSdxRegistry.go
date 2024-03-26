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

type NetScalerSdxRegistry struct {
	// Environments []Environment `json:"environments" yaml:"environments" mapstructure:"environments" secure:"true"`
}

func (r NetScalerSdxRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerSdxRegistry{},
		Encrypted: SecureNetScalerSdxRegistry{},
	}
}

type SecureNetScalerSdxRegistry struct {
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureNetScalerSdxRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureNetScalerSdxRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerSdxRegistry{},
		Encrypted: SecureNetScalerSdxRegistry{},
	}
}
