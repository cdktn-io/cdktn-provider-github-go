// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package teammembers


type TeamMembersMembers struct {
	// User to add to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/team_members#username TeamMembers#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Role to grant the user within the team; must be one of `member` or `maintainer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.13.0/docs/resources/team_members#role TeamMembers#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

