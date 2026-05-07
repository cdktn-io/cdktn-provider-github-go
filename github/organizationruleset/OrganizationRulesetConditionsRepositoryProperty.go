// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetConditionsRepositoryProperty struct {
	// The repository properties and values to exclude. The ruleset will not apply if any of these properties match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.12.1/docs/resources/organization_ruleset#exclude OrganizationRuleset#exclude}
	Exclude interface{} `field:"optional" json:"exclude" yaml:"exclude"`
	// The repository properties and values to include. All of these properties must match for the condition to pass.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.12.1/docs/resources/organization_ruleset#include OrganizationRuleset#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
}

