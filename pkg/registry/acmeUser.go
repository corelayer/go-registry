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

func NewAcmeUser(name string, email string, eab AcmeExternalAccountBinding) AcmeUser {
	return AcmeUser{
		Name:                   name,
		Email:                  email,
		ExternalAccountBinding: eab,
	}
}

type AcmeUser struct {
	Name                   string                     `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Email                  string                     `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty" secure:"true"`
	ExternalAccountBinding AcmeExternalAccountBinding `json:"eab,omitempty" yaml:"eab,omitempty" mapstructure:"eab,omitempty" secure:"true"`
}

func (u AcmeUser) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeUser{},
		Encrypted: SecureAcmeUser{},
	}
}

type SecureAcmeUser struct {
	Name                   string                           `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Email                  string                           `json:"email,omitempty" yaml:"email,omitempty" mapstructure:"email,omitempty" secure:"true"`
	ExternalAccountBinding SecureAcmeExternalAccountBinding `json:"eab,omitempty" yaml:"eab,omitempty" mapstructure:"eab,omitempty" secure:"true"`
	CryptoParams           cryptostruct.CryptoParams        `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureAcmeUser) Decrypt(key string) (AcmeUser, error) {
	var (
		err       error
		decrypted any
	)
	decrypter := cryptostruct.NewDecrypter(hex.EncodeToString([]byte(key)), s.GetTransformConfig())
	if decrypted, err = decrypter.Transform(s); err != nil {
		return AcmeUser{}, err
	}
	return decrypted.(AcmeUser), nil
}

func (s SecureAcmeUser) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureAcmeUser) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeUser{},
		Encrypted: SecureAcmeUser{},
	}
}
