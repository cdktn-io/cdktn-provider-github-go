// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredStatusChecksRequiredCheck struct {
	// The status check context name that must be present on the commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#context RepositoryRuleset#context}
	Context *string `field:"required" json:"context" yaml:"context"`
	// The optional integration ID that this status check must originate from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.11.1/docs/resources/repository_ruleset#integration_id RepositoryRuleset#integration_id}
	IntegrationId *float64 `field:"optional" json:"integrationId" yaml:"integrationId"`
}

