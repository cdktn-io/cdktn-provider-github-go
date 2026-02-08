// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesPullRequestRequiredReviewersReviewer struct {
	// The ID of the reviewer that must review.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/organization_ruleset#id OrganizationRuleset#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// The type of reviewer. Currently only `Team` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/organization_ruleset#type OrganizationRuleset#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

