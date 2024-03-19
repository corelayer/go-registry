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

type AcmeExternalAccountBinding struct {
	Kid  string `json:"kid,omitempty" yaml:"kid,omitempty" mapstructure:"kid,omitempty" secure:"true"`
	Hmac string `json:"hmac,omitempty" yaml:"hmac,omitempty" mapstructure:"hmac,omitempty" secure:"true"`
}

func (e AcmeExternalAccountBinding) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeExternalAccountBinding{},
		Encrypted: SecureAcmeExternalAccountBinding{},
	}
}

type SecureAcmeExternalAccountBinding struct {
	Kid          string                    `json:"kid,omitempty" yaml:"kid,omitempty" mapstructure:"kid,omitempty" secure:"true"`
	Hmac         string                    `json:"hmac,omitempty" yaml:"hmac,omitempty" mapstructure:"hmac,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureAcmeExternalAccountBinding) Decrypt(key string) (AcmeExternalAccountBinding, error) {
	var (
		err       error
		decrypted any
	)
	decrypter := cryptostruct.NewDecrypter(hex.EncodeToString([]byte(key)), s.GetTransformConfig())
	if decrypted, err = decrypter.Transform(s); err != nil {
		return AcmeExternalAccountBinding{}, err
	}
	return decrypted.(AcmeExternalAccountBinding), nil
}

func (s SecureAcmeExternalAccountBinding) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureAcmeExternalAccountBinding) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeExternalAccountBinding{},
		Encrypted: SecureAcmeExternalAccountBinding{},
	}
}
