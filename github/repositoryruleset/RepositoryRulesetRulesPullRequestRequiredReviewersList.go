// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-github-go/github/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-github-go/github/v15/repositoryruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetRulesPullRequestRequiredReviewersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryRulesetRulesPullRequestRequiredReviewersList
type jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRepositoryRulesetRulesPullRequestRequiredReviewersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) RepositoryRulesetRulesPullRequestRequiredReviewersList {
	_init_.Initialize()

	if err := validateNewRepositoryRulesetRulesPullRequestRequiredReviewersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList{}

	_jsii_.Create(
		"@cdktn/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestRequiredReviewersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRepositoryRulesetRulesPullRequestRequiredReviewersList_Override(r RepositoryRulesetRulesPullRequestRequiredReviewersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestRequiredReviewersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := r.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		r,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) Get(index *float64) RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) Resolve(context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

