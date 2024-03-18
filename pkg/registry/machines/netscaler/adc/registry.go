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

import (
	"fmt"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type Registry struct {
	Environments []Environment `json:"environments,omitempty" yaml:"environments,omitempty" mapstructure:"environments,omitempty" secure:"true"`
}

func (r Registry) GetEnvironmentByName(name string) (Environment, error) {
	for _, e := range r.Environments {
		if e.Name == name {
			return e, nil
		}
	}
	return Environment{}, fmt.Errorf("could not find environment %s", name)
}

func (r Registry) GetEnvironmentNames() []string {
	var (
		output []string
	)
	output = make([]string, len(r.Environments))
	for i := 0; i < len(r.Environments); i++ {
		output[i] = r.Environments[i].Name
	}

	return output
}

func (r Registry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

type SecureRegistry struct {
	Environments []SecureEnvironment       `json:"environments,omitempty" yaml:"environments,omitempty" mapstructure:"environments,omitempty" secure:"true"`
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
