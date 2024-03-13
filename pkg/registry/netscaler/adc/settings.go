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

type Settings struct {
	UseSsl                    bool   `json:"useSsl,omitempty" yaml:"useSsl,omitempty" mapstructure:"useSsl,omitempty"`
	Timeout                   int    `json:"timeout,omitempty" yaml:"timeout,omitempty" mapstructure:"timeout,omitempty"`
	UserAgent                 string `json:"userAgent,omitempty" yaml:"userAgent,omitempty" mapstructure:"userAgent,omitempty"`
	ValidateServerCertificate bool   `json:"validateServerCertificate,omitempty" yaml:"validateServerCertificate,omitempty" mapstructure:"validateServerCertificate,omitempty"`
	LogTlsSecrets             bool   `json:"logTlsSecrets,omitempty" yaml:"logTlsSecrets,omitempty" mapstructure:"logTlsSecrets,omitempty"`
	LogTlsSecretsDestination  string `json:"logTlsSecretsDestination,omitempty" yaml:"logTlsSecretsDestination,omitempty" mapstructure:"logTlsSecretsDestination,omitempty"`
	AutoLogin                 bool   `json:"autoLogin,omitempty" yaml:"autoLogin,omitempty" mapstructure:"autoLogin,omitempty"`
}
