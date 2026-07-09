// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetBypassActors struct {
	// The type of actor that can bypass a ruleset.
	//
	// Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`, `DeployKey`, `EnterpriseOwner`, or `User`. See https://docs.github.com/en/rest/repos/rules for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_ruleset#actor_type RepositoryRuleset#actor_type}
	ActorType *string `field:"required" json:"actorType" yaml:"actorType"`
	// When the specified actor can bypass the ruleset.
	//
	// pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_ruleset#bypass_mode RepositoryRuleset#bypass_mode}
	BypassMode *string `field:"required" json:"bypassMode" yaml:"bypassMode"`
	// The ID of the actor that can bypass a ruleset.
	//
	// If `actor_type` is `Integration`, `actor_id` is a GitHub App ID. When `actor_type` is `User`, this should be set to the numeric GitHub user ID. Some actor types such as `OrganizationAdmin`, `EnterpriseOwner`, and `DeployKey` do not have an ID — this argument should not be set in those cases as the GitHub API will ignore it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/repository_ruleset#actor_id RepositoryRuleset#actor_id}
	ActorId *float64 `field:"optional" json:"actorId" yaml:"actorId"`
}

