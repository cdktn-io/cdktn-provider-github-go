// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetBypassActors struct {
	// The type of actor that can bypass a ruleset.
	//
	// Can be one of: `Integration`, `OrganizationAdmin`, `RepositoryRole`, `Team`, `DeployKey`, or `EnterpriseOwner`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#actor_type OrganizationRuleset#actor_type}
	ActorType *string `field:"required" json:"actorType" yaml:"actorType"`
	// When the specified actor can bypass the ruleset.
	//
	// pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#bypass_mode OrganizationRuleset#bypass_mode}
	BypassMode *string `field:"required" json:"bypassMode" yaml:"bypassMode"`
	// The ID of the actor that can bypass a ruleset.
	//
	// Must be omitted for ID-less actor types: `OrganizationAdmin`, `EnterpriseOwner`, and `DeployKey` — the GitHub API does not use an ID for these types and will ignore any value set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/organization_ruleset#actor_id OrganizationRuleset#actor_id}
	ActorId *float64 `field:"optional" json:"actorId" yaml:"actorId"`
}

