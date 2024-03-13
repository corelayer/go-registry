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
	"fmt"

	"github.com/corelayer/go-cryptostruct/pkg/cryptostruct"
)

type Registry struct {
	Services  []Service  `json:"services,omitempty" yaml:"services,omitempty" mapstructure:"services,omitempty" secure:"true"`
	Users     []User     `json:"users,omitempty" yaml:"users,omitempty" mapstructure:"users,omitempty" secure:"true"`
	Providers []Provider `json:"providers,omitempty" yaml:"providers,omitempty" mapstructure:"providers,omitempty" secure:"true"`
}

func (r Registry) GetProviderByName(name string) (Provider, error) {
	for _, p := range r.Providers {
		if p.Name == name {
			return p, nil
		}
	}
	return Provider{}, fmt.Errorf("could not find provider %s", name)
}

func (r Registry) GetProviderNames() []string {
	var (
		output []string
	)
	output = make([]string, len(r.Providers))
	for i := 0; i < len(r.Providers); i++ {
		output[i] = r.Providers[i].Name
	}

	return output
}

func (r Registry) GetServiceByName(name string) (Service, error) {
	for _, s := range r.Services {
		if s.Name == name {
			return s, nil
		}
	}
	return Service{}, fmt.Errorf("could not find service %s", name)
}

func (r Registry) GetServiceNames() []string {
	var (
		output []string
	)
	output = make([]string, len(r.Services))
	for i := 0; i < len(r.Services); i++ {
		output[i] = r.Services[i].Name
	}

	return output
}

func (r Registry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

func (r Registry) GetUserByName(name string) (User, error) {
	for _, u := range r.Users {
		if u.Name == name {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("could not find user %s", name)
}

func (r Registry) GetUserNames() []string {
	var (
		output []string
	)
	output = make([]string, len(r.Users))
	for i := 0; i < len(r.Users); i++ {
		output[i] = r.Users[i].Name
	}

	return output
}

type SecureRegistry struct {
	Services     []SecureService           `json:"services,omitempty" yaml:"services,omitempty" mapstructure:"services,omitempty" secure:"true"`
	Users        []SecureUser              `json:"users,omitempty" yaml:"users,omitempty" mapstructure:"users,omitempty" secure:"true"`
	Providers    []SecureProvider          `json:"providers,omitempty" yaml:"providers,omitempty" mapstructure:"providers,omitempty" secure:"true"`
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
