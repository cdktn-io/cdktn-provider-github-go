// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesCopilotCodeReview struct {
	// Copilot automatically reviews draft pull requests before they are marked as ready for review. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/repository_ruleset#review_draft_pull_requests RepositoryRuleset#review_draft_pull_requests}
	ReviewDraftPullRequests interface{} `field:"optional" json:"reviewDraftPullRequests" yaml:"reviewDraftPullRequests"`
	// Copilot automatically reviews each new push to the pull request. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.0/docs/resources/repository_ruleset#review_on_push RepositoryRuleset#review_on_push}
	ReviewOnPush interface{} `field:"optional" json:"reviewOnPush" yaml:"reviewOnPush"`
}

