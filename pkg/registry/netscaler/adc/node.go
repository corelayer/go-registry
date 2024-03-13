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

package adc

import "github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

type Node struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address string `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
}

func (n Node) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Node{},
		Encrypted: SecureNode{},
	}
}

type SecureNode struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address      string                    `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureNode) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Node{},
		Encrypted: SecureNode{},
	}
}

func (s SecureNode) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}
