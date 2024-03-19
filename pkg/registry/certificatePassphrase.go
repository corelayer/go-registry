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
	"encoding/hex"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type CertificatePassphrase struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Value string `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty" secure:"true"`
}

func (c CertificatePassphrase) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: CertificatePassphrase{},
		Encrypted: SecureCertificatePassphrase{},
	}
}

type SecureCertificatePassphrase struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Value        string                    `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureCertificatePassphrase) Decrypt(key string) (CertificatePassphrase, error) {
	var (
		err       error
		decrypted any
	)
	decrypter := cryptostruct.NewDecrypter(hex.EncodeToString([]byte(key)), s.GetTransformConfig())
	if decrypted, err = decrypter.Transform(s); err != nil {
		return CertificatePassphrase{}, err
	}
	return decrypted.(CertificatePassphrase), nil
}

func (s SecureCertificatePassphrase) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureCertificatePassphrase) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: CertificatePassphrase{},
		Encrypted: SecureCertificatePassphrase{},
	}
}
