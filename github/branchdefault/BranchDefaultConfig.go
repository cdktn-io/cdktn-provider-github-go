// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package branchdefault

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BranchDefaultConfig struct {
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
	// The name of the branch to set as the default (e.g. 'main').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#branch BranchDefault#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The name of the GitHub repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#repository BranchDefault#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The ETag header for the repository API response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#etag BranchDefault#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#id BranchDefault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If `true` rename the existing branch when the `branch` input is changed. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#rename BranchDefault#rename}
	Rename interface{} `field:"optional" json:"rename" yaml:"rename"`
	// If `true`, poll until GitHub propagates the renamed default branch before proceeding.
	//
	// Only has effect when `rename` is also `true`. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/branch_default#wait_for_rename BranchDefault#wait_for_rename}
	WaitForRename interface{} `field:"optional" json:"waitForRename" yaml:"waitForRename"`
}

