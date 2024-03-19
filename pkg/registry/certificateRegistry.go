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

type CertificateRegistry struct {
	Acme        AcmeRegistry            `json:"acme,omitempty" yaml:"acme,omitempty" mapstructure:"acme,omitempty" secure:"true"`
	Passphrases []CertificatePassphrase `json:"passphrases,omitempty" yaml:"passphrases,omitempty" mapstructure:"passphrases,omitempty" secure:"true"`
}

func (r CertificateRegistry) GetPassphraseByName(name string) (CertificatePassphrase, error) {
	for _, p := range r.Passphrases {
		if p.Name == name {
			return p, nil
		}
	}
	return CertificatePassphrase{}, NewItemNotFoundError("passphrase", name)
}

func (r CertificateRegistry) GetPassphraseNames() []string {
	names := make([]string, len(r.Passphrases))
	for _, p := range r.Passphrases {
		names = append(names, p.Name)
	}
	return names
}

func (r CertificateRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: CertificateRegistry{},
		Encrypted: SecureCertificateRegistry{},
	}
}

type SecureCertificateRegistry struct {
	Acme         SecureAcmeRegistry            `json:"acme,omitempty" yaml:"acme,omitempty" mapstructure:"acme,omitempty" secure:"true"`
	Passphrases  []SecureCertificatePassphrase `json:"passphrases,omitempty" yaml:"passphrases,omitempty" mapstructure:"passphrases,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams     `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureCertificateRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (r SecureCertificateRegistry) GetPassphraseByName(name string) (SecureCertificatePassphrase, error) {
	for _, p := range r.Passphrases {
		if p.Name == name {
			return p, nil
		}
	}
	return SecureCertificatePassphrase{}, NewItemNotFoundError("passphrase", name)
}

func (r SecureCertificateRegistry) GetPassphraseNames() []string {
	names := make([]string, len(r.Passphrases))
	for _, p := range r.Passphrases {
		names = append(names, p.Name)
	}
	return names
}

func (s SecureCertificateRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: CertificateRegistry{},
		Encrypted: SecureCertificateRegistry{},
	}
}
