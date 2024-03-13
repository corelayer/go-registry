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

import (
	"log/slog"
	"os"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type Provider struct {
	Name      string     `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Type      string     `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty" secure:"false"`
	Challenge string     `json:"challenge,omitempty" yaml:"challenge,omitempty" mapstructure:"challenge,omitempty" secure:"false"`
	Variables []Variable `json:"variables,omitempty" yaml:"variables,omitempty" mapstructure:"variables,omitempty" secure:"true"`
}

func (p Provider) ApplyEnvironmentVariables() error {
	var err error
	slog.Debug("applying provider parameters", "name", p.Name)
	for _, v := range p.Variables {
		slog.Debug("applying provider parameter", "name", p.Name, "variable", v.Key)
		err = os.Setenv(v.Key, v.Value)
		if err != nil {
			return err
		}
	}
	slog.Debug("applying provider parameters completed", "name", p.Name)
	return nil
}

func (p Provider) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Provider{},
		Encrypted: SecureProvider{},
	}
}

func (p Provider) ResetEnvironmentVariables() error {
	var err error
	slog.Debug("resetting provider parameters", "name", p.Name)
	for _, v := range p.Variables {
		slog.Debug("resetting provider parameter", "name", p.Name, "variable", v.Key)
		err = os.Unsetenv(v.Key)
		if err != nil {
			return err
		}
	}
	slog.Debug("resetting provider parameters completed", "name", p.Name)
	return nil
}

type SecureProvider struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Type         string                    `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty" secure:"false"`
	Challenge    string                    `json:"challenge,omitempty" yaml:"challenge,omitempty" mapstructure:"challenge,omitempty" secure:"false"`
	Variables    []SecureVariable          `json:"variables,omitempty" yaml:"variables,omitempty" mapstructure:"variables,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureProvider) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Provider{},
		Encrypted: SecureProvider{},
	}
}

func (s SecureProvider) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
