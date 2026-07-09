// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type GithubProviderAppAuth struct {
	// The GitHub App's identifier. This can also be set by the `GITHUB_APP_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#id GithubProvider#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The GitHub App's installation identifier. This can also be set by the `GITHUB_APP_INSTALLATION_ID` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#installation_id GithubProvider#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// The GitHub App's PEM file content;
	//
	// `\n` can be used for newlines. This can also be set by the `GITHUB_APP_PEM_FILE` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs#pem_file GithubProvider#pem_file}
	PemFile *string `field:"required" json:"pemFile" yaml:"pemFile"`
}

