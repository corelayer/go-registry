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
	"fmt"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

	"github.com/corelayer/go-registry/pkg/registry/certificates/acme"
)

type Registry struct {
	Acme        acme.Registry `json:"acme,omitempty" yaml:"acme,omitempty" mapstructure:"acme,omitempty" secure:"true"`
	Passphrases []Passphrase  `json:"passphrases,omitempty" yaml:"passphrases,omitempty" mapstructure:"passphrases,omitempty" secure:"true"`
}

func (r Registry) GetPassphraseByName(name string) (Passphrase, error) {
	for _, p := range r.Passphrases {
		if p.Name == name {
			return p, nil
		}
	}
	return Passphrase{}, fmt.Errorf("could not find passphrase %s", name)
}

func (r Registry) GetPassphraseNames() []string {
	names := make([]string, len(r.Passphrases))
	for _, p := range r.Passphrases {
		names = append(names, p.Name)
	}
	return names
}

func (r Registry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

type SecureRegistry struct {
	Acme         acme.SecureRegistry       `json:"acme,omitempty" yaml:"acme,omitempty" mapstructure:"acme,omitempty" secure:"true"`
	Passphrases  []SecurePassphrase        `json:"passphrases,omitempty" yaml:"passphrases,omitempty" mapstructure:"passphrases,omitempty" secure:"true"`
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
