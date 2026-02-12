// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesCopilotCodeReview struct {
	// Copilot automatically reviews draft pull requests before they are marked as ready for review. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#review_draft_pull_requests OrganizationRuleset#review_draft_pull_requests}
	ReviewDraftPullRequests interface{} `field:"optional" json:"reviewDraftPullRequests" yaml:"reviewDraftPullRequests"`
	// Copilot automatically reviews each new push to the pull request. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/organization_ruleset#review_on_push OrganizationRuleset#review_on_push}
	ReviewOnPush interface{} `field:"optional" json:"reviewOnPush" yaml:"reviewOnPush"`
}

