// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-github-go/github/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-github-go/github/v15/repositoryruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FilePatterns() *[]*string
	SetFilePatterns(val *[]*string)
	FilePatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MinimumApprovals() *float64
	SetMinimumApprovals(val *float64)
	MinimumApprovalsInput() *float64
	Reviewer() RepositoryRulesetRulesPullRequestRequiredReviewersReviewerOutputReference
	ReviewerInput() *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutReviewer(value *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference
type jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) FilePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) FilePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) MinimumApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) MinimumApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) Reviewer() RepositoryRulesetRulesPullRequestRequiredReviewersReviewerOutputReference {
	var returns RepositoryRulesetRulesPullRequestRequiredReviewersReviewerOutputReference
	_jsii_.Get(
		j,
		"reviewer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) ReviewerInput() *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer {
	var returns *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer
	_jsii_.Get(
		j,
		"reviewerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRepositoryRulesetRulesPullRequestRequiredReviewersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference {
	_init_.Initialize()

	if err := validateNewRepositoryRulesetRulesPullRequestRequiredReviewersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRepositoryRulesetRulesPullRequestRequiredReviewersOutputReference_Override(r RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetFilePatterns(val *[]*string) {
	if err := j.validateSetFilePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePatterns",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetMinimumApprovals(val *float64) {
	if err := j.validateSetMinimumApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumApprovals",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) PutReviewer(value *RepositoryRulesetRulesPullRequestRequiredReviewersReviewer) {
	if err := r.validatePutReviewerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putReviewer",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestRequiredReviewersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

