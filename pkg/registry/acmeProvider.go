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
	"os"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type AcmeProvider struct {
	Name      string         `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Type      string         `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty" secure:"false"`
	Challenge string         `json:"challenge,omitempty" yaml:"challenge,omitempty" mapstructure:"challenge,omitempty" secure:"false"`
	Variables []AcmeVariable `json:"variables,omitempty" yaml:"variables,omitempty" mapstructure:"variables,omitempty" secure:"true"`
}

func (p AcmeProvider) ApplyEnvironmentVariables() error {
	for _, v := range p.Variables {
		if err := os.Setenv(v.Key, v.Value); err != nil {
			return err
		}
	}
	return nil
}

func (p AcmeProvider) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeProvider{},
		Encrypted: SecureAcmeProvider{},
	}
}

func (p AcmeProvider) ResetEnvironmentVariables() error {
	for _, v := range p.Variables {
		if err := os.Unsetenv(v.Key); err != nil {
			return err
		}
	}
	return nil
}

type SecureAcmeProvider struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Type         string                    `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty" secure:"false"`
	Challenge    string                    `json:"challenge,omitempty" yaml:"challenge,omitempty" mapstructure:"challenge,omitempty" secure:"false"`
	Variables    []SecureAcmeVariable      `json:"variables,omitempty" yaml:"variables,omitempty" mapstructure:"variables,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (p SecureAcmeProvider) ApplyEnvironmentVariables(key string) error {
	for _, sv := range p.Variables {
		var (
			err error
			v   AcmeVariable
		)
		if v, err = sv.Decrypt(key); err != nil {
			return err
		}

		if err = os.Setenv(v.Key, v.Value); err != nil {
			return err
		}
	}
	return nil
}

func (s SecureAcmeProvider) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureAcmeProvider) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeProvider{},
		Encrypted: SecureAcmeProvider{},
	}
}

func (p SecureAcmeProvider) ResetEnvironmentVariables() error {
	for _, v := range p.Variables {
		if err := os.Unsetenv(v.Key); err != nil {
			return err
		}
	}
	return nil
}
