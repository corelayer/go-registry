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

type AcmeRegistry struct {
	Services  []AcmeService  `json:"services,omitempty" yaml:"services,omitempty" mapstructure:"services,omitempty" secure:"true"`
	Users     []AcmeUser     `json:"users,omitempty" yaml:"users,omitempty" mapstructure:"users,omitempty" secure:"true"`
	Providers []AcmeProvider `json:"providers,omitempty" yaml:"providers,omitempty" mapstructure:"providers,omitempty" secure:"true"`
}

func (r AcmeRegistry) GetProviderByName(name string) (AcmeProvider, error) {
	for _, p := range r.Providers {
		if p.Name == name {
			return p, nil
		}
	}
	return AcmeProvider{}, NewItemNotFoundError("acme provider", name)
}

func (r AcmeRegistry) GetProviderNames() []string {
	names := make([]string, len(r.Providers))
	for i := 0; i < len(r.Providers); i++ {
		names[i] = r.Providers[i].Name
	}

	return names
}

func (r AcmeRegistry) GetServiceByName(name string) (AcmeService, error) {
	for _, s := range r.Services {
		if s.Name == name {
			return s, nil
		}
	}
	return AcmeService{}, NewItemNotFoundError("acme service", name)
}

func (r AcmeRegistry) GetServiceNames() []string {
	names := make([]string, len(r.Services))
	for i := 0; i < len(r.Services); i++ {
		names[i] = r.Services[i].Name
	}

	return names
}

func (r AcmeRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeRegistry{},
		Encrypted: SecureAcmeRegistry{},
	}
}

func (r AcmeRegistry) GetUserByName(name string) (AcmeUser, error) {
	for _, u := range r.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return AcmeUser{}, NewItemNotFoundError("acme user", name)
}

func (r AcmeRegistry) GetUserNames() []string {
	names := make([]string, len(r.Users))
	for i := 0; i < len(r.Users); i++ {
		names[i] = r.Users[i].Name
	}

	return names
}

type SecureAcmeRegistry struct {
	Services     []SecureAcmeService       `json:"services,omitempty" yaml:"services,omitempty" mapstructure:"services,omitempty" secure:"true"`
	Users        []SecureAcmeUser          `json:"users,omitempty" yaml:"users,omitempty" mapstructure:"users,omitempty" secure:"true"`
	Providers    []SecureAcmeProvider      `json:"providers,omitempty" yaml:"providers,omitempty" mapstructure:"providers,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureAcmeRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureAcmeRegistry) GetProviderByName(name string) (SecureAcmeProvider, error) {
	for _, p := range s.Providers {
		if p.Name == name {
			return p, nil
		}
	}
	return SecureAcmeProvider{}, NewItemNotFoundError("acme provider", name)
}

func (s SecureAcmeRegistry) GetProviderNames() []string {
	names := make([]string, len(s.Providers))
	for i := 0; i < len(s.Providers); i++ {
		names[i] = s.Providers[i].Name
	}

	return names
}

func (s SecureAcmeRegistry) GetServiceByName(name string) (SecureAcmeService, error) {
	for _, s := range s.Services {
		if s.Name == name {
			return s, nil
		}
	}
	return SecureAcmeService{}, NewItemNotFoundError("acme service", name)
}

func (s SecureAcmeRegistry) GetServiceNames() []string {
	names := make([]string, len(s.Services))
	for i := 0; i < len(s.Services); i++ {
		names[i] = s.Services[i].Name
	}

	return names
}

func (s SecureAcmeRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: AcmeRegistry{},
		Encrypted: SecureAcmeRegistry{},
	}
}

func (s SecureAcmeRegistry) GetUserByName(name string) (SecureAcmeUser, error) {
	for _, u := range s.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return SecureAcmeUser{}, NewItemNotFoundError("acme user", name)
}

func (s SecureAcmeRegistry) GetUserNames() []string {
	names := make([]string, len(s.Users))
	for i := 0; i < len(s.Users); i++ {
		names[i] = s.Users[i].Name
	}

	return names
}
