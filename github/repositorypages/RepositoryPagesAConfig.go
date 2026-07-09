// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositorypages

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RepositoryPagesAConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The repository name to configure GitHub Pages for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#repository RepositoryPagesA#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The type of GitHub Pages site to build. Can be 'legacy' or 'workflow'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#build_type RepositoryPagesA#build_type}
	BuildType *string `field:"optional" json:"buildType" yaml:"buildType"`
	// The custom domain for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#cname RepositoryPagesA#cname}
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
	// Whether the rendered GitHub Pages site will only be served over HTTPS. Requires 'cname' to be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#https_enforced RepositoryPagesA#https_enforced}
	HttpsEnforced interface{} `field:"optional" json:"httpsEnforced" yaml:"httpsEnforced"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#id RepositoryPagesA#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the GitHub Pages site is publicly visible.
	//
	// If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#public RepositoryPagesA#public}
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_pages#source RepositoryPagesA#source}
	Source *RepositoryPagesSourceA `field:"optional" json:"source" yaml:"source"`
}

