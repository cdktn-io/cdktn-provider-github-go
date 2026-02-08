// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationvariablerepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationVariableRepositoryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The repository ID that can access the organization variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/actions_organization_variable_repository#repository_id ActionsOrganizationVariableRepository#repository_id}
	RepositoryId *float64 `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Name of the existing variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/actions_organization_variable_repository#variable_name ActionsOrganizationVariableRepository#variable_name}
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/actions_organization_variable_repository#id ActionsOrganizationVariableRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

