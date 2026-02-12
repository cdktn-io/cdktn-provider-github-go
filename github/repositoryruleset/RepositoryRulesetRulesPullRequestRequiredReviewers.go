// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesPullRequestRequiredReviewers struct {
	// File patterns (fnmatch syntax) that this reviewer must approve.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#file_patterns RepositoryRuleset#file_patterns}
	FilePatterns *[]*string `field:"required" json:"filePatterns" yaml:"filePatterns"`
	// Minimum number of approvals required from this reviewer. Set to 0 to make approval optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#minimum_approvals RepositoryRuleset#minimum_approvals}
	MinimumApprovals *float64 `field:"required" json:"minimumApprovals" yaml:"minimumApprovals"`
	// reviewer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#reviewer RepositoryRuleset#reviewer}
	Reviewer *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer `field:"required" json:"reviewer" yaml:"reviewer"`
}

