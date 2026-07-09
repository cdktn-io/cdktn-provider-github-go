// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetConditionsRepositoryPropertyInclude struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#name OrganizationRuleset#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#property_values OrganizationRuleset#property_values}.
	PropertyValues *[]*string `field:"optional" json:"propertyValues" yaml:"propertyValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#source OrganizationRuleset#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

