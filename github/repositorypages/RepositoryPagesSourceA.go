// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositorypages


type RepositoryPagesSourceA struct {
	// The repository branch used to publish the site's source files. (i.e. 'main' or 'gh-pages').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.12.1/docs/resources/repository_pages#branch RepositoryPagesA#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The repository directory from which the site publishes (Default: '/').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.12.1/docs/resources/repository_pages#path RepositoryPagesA#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

