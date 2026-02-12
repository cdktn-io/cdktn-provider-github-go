// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetConfig struct {
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
	// The enforcement level of the ruleset.
	//
	// `evaluate` allows admins to test rules before enforcing them. Possible values are `disabled`, `active`, and `evaluate`. Note: `evaluate` is only available for Enterprise plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#enforcement OrganizationRuleset#enforcement}
	Enforcement *string `field:"required" json:"enforcement" yaml:"enforcement"`
	// The name of the ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#name OrganizationRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#rules OrganizationRuleset#rules}
	Rules *OrganizationRulesetRules `field:"required" json:"rules" yaml:"rules"`
	// The target of the ruleset. Possible values are branch, tag and push.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#target OrganizationRuleset#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// bypass_actors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#bypass_actors OrganizationRuleset#bypass_actors}
	BypassActors interface{} `field:"optional" json:"bypassActors" yaml:"bypassActors"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#conditions OrganizationRuleset#conditions}
	Conditions *OrganizationRulesetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#id OrganizationRuleset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

