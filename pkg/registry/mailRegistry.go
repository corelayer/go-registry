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

func NewMailRegistry() MailRegistry {
	return MailRegistry{
		SmtpServers: make([]SmtpServer, 0),
	}
}

type MailRegistry struct {
	SmtpServers []SmtpServer `json:"smtpServers,omitempty" yaml:"smtpServers,omitempty" mapstructure:"smtpServers,omitempty" secure:"true"`
}

func (r MailRegistry) GetSmtpServerByName(name string) (SmtpServer, error) {
	for _, s := range r.SmtpServers {
		if s.Name == name {
			return s, nil
		}
	}
	return SmtpServer{}, NewItemNotFoundError("smtp server", name)
}

func (r MailRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: MailRegistry{},
		Encrypted: SecureMailRegistry{},
	}
}

type SecureMailRegistry struct {
	SmtpServers  []SecureSmtpServer        `json:"smtpServers,omitempty" yaml:"smtpServers,omitempty" mapstructure:"smtpServers,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureMailRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (r SecureMailRegistry) GetSmtpServerByName(name string) (SecureSmtpServer, error) {
	for _, s := range r.SmtpServers {
		if s.Name == name {
			return s, nil
		}
	}
	return SecureSmtpServer{}, NewItemNotFoundError("smtp server", name)
}

func (s SecureMailRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: MailRegistry{},
		Encrypted: SecureMailRegistry{},
	}
}
