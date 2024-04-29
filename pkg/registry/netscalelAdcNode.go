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

func NewNetScalerAdcNode(name string, address string) NetScalerAdcNode {
	return NetScalerAdcNode{
		Name:    name,
		Address: address,
	}
}

type NetScalerAdcNode struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address string `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
}

func (n NetScalerAdcNode) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcNode{},
		Encrypted: SecureNetScalerAdcNode{},
	}
}

type SecureNetScalerAdcNode struct {
	Name         string                    `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" secure:"false"`
	Address      string                    `json:"address,omitempty" yaml:"address,omitempty" mapstructure:"address,omitempty" secure:"true"`
	CryptoParams cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureNetScalerAdcNode) Decrypt(key string) (NetScalerAdcNode, error) {
	var (
		err       error
		decrypted any
	)
	decrypter := cryptostruct.NewDecrypter(hex.EncodeToString([]byte(key)), s.GetTransformConfig())
	if decrypted, err = decrypter.Transform(s); err != nil {
		return NetScalerAdcNode{}, err
	}
	return decrypted.(NetScalerAdcNode), nil
}

func (s SecureNetScalerAdcNode) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureNetScalerAdcNode) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: NetScalerAdcNode{},
		Encrypted: SecureNetScalerAdcNode{},
	}
}
