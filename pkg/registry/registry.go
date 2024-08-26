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

import "github.com/corelayer/go-cryptostruct/pkg/cryptostruct"

func NewEmptyRegistry() Registry {
	return Registry{
		Organizations: []Organization{NewOrganization("")},
	}
}

func NewExampleRegistry() Registry {
	return Registry{
		Organizations: []Organization{
			{
				Name: "organization name",
				Registry: OrganizationRegistry{
					Machines: MachinesRegistry{
						NetScaler: NetScalerRegistry{
							Adc: NetScalerAdcRegistry{
								Environments: []NetScalerAdcEnvironment{
									{
										Name: "netscaler adc environment name",
										Management: NetScalerAdcNode{
											Name:    "netscaler adc snip name",
											Address: "netscaler adc snip ip address | fqdn",
										},
										Nodes: []NetScalerAdcNode{
											{
												Name:    "netscaler adc nsip name",
												Address: "netscaler adc nsip ip address | fqdn",
											},
										},
										Credentials: []NetScalerAdcCredential{
											{
												Name:     "netscaler adc credential name",
												Username: "netscaler adc credential username",
												Password: "netscaler adc credential password",
											},
										},
										Settings: NetScalerAdcSettings{
											UseSsl:                    false,
											Timeout:                   30,
											UserAgent:                 "go-cts",
											ValidateServerCertificate: true,
											LogTlsSecrets:             false,
											LogTlsSecretsDestination:  "",
											AutoLogin:                 false,
										},
									},
								},
							},
						},
					},
					Certificates: CertificateRegistry{
						Acme: AcmeRegistry{
							Services: []AcmeService{
								{
									Name: "acme service name",
									Url:  "acme service url",
								},
							},
							Users: []AcmeUser{
								{
									Name:  "acme user name",
									Email: "acme user email",
									ExternalAccountBinding: AcmeExternalAccountBinding{
										Kid:  "acme user kid value",
										Hmac: "acme user hmac value",
									},
								},
							},
							Providers: []AcmeProvider{
								{
									Name:      "acme provider name",
									Type:      "acme provider type",
									Challenge: "acme challenge type http-01 | dns-01",
									Variables: []AcmeVariable{
										{
											Key:   "acme provider variable name",
											Value: "acme provider variable value",
										},
									},
								},
							},
						},
						Passphrases: []CertificatePassphrase{
							{
								Name:  "passphrase name",
								Value: "passphrase value",
							},
						},
					},
				},
			},
		}}
}

type Registry struct {
	Organizations []Organization `json:"organizations,omitempty" yaml:"organizations,omitempty" mapstructure:"organizations,omitempty" secure:"true"`
}

func (r Registry) GetOrganizationByName(name string) (Organization, error) {
	for _, o := range r.Organizations {
		if o.Name == name {
			return o, nil
		}
	}
	return Organization{}, NewItemNotFoundError("organization", name)
}

func (r Registry) GetOrganizationNames() []string {
	names := make([]string, len(r.Organizations))
	for i := 0; i < len(r.Organizations); i++ {
		names[i] = r.Organizations[i].Name
	}

	return names
}

func (o Registry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}

type SecureRegistry struct {
	Organizations []SecureOrganization      `json:"organizations,omitempty" yaml:"organizations,omitempty" mapstructure:"organizations,omitempty" secure:"true"`
	CryptoParams  cryptostruct.CryptoParams `json:"cryptoParams" yaml:"cryptoParams" mapstructure:"cryptoParams"`
}

func (s SecureRegistry) GetCryptoParams() cryptostruct.CryptoParams {
	return s.CryptoParams
}

func (s SecureRegistry) GetOrganizationByName(name string) (SecureOrganization, error) {
	for _, o := range s.Organizations {
		if o.Name == name {
			return o, nil
		}
	}
	return SecureOrganization{}, NewItemNotFoundError("organization", name)
}

func (s SecureRegistry) GetOrganizationNames() []string {
	names := make([]string, len(s.Organizations))
	for i := 0; i < len(s.Organizations); i++ {
		names[i] = s.Organizations[i].Name
	}

	return names
}

func (s SecureRegistry) GetTransformConfig() cryptostruct.TransformConfig {
	return cryptostruct.TransformConfig{
		Decrypted: Registry{},
		Encrypted: SecureRegistry{},
	}
}
