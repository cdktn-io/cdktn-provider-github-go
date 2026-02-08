// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-github-go/github/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-github-go/github/v15/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference interface {
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
	Reviewer() OrganizationRulesetRulesPullRequestRequiredReviewersReviewerOutputReference
	ReviewerInput() *OrganizationRulesetRulesPullRequestRequiredReviewersReviewer
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
	PutReviewer(value *OrganizationRulesetRulesPullRequestRequiredReviewersReviewer)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference
type jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) FilePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) FilePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) MinimumApprovals() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) MinimumApprovalsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumApprovalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) Reviewer() OrganizationRulesetRulesPullRequestRequiredReviewersReviewerOutputReference {
	var returns OrganizationRulesetRulesPullRequestRequiredReviewersReviewerOutputReference
	_jsii_.Get(
		j,
		"reviewer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) ReviewerInput() *OrganizationRulesetRulesPullRequestRequiredReviewersReviewer {
	var returns *OrganizationRulesetRulesPullRequestRequiredReviewersReviewer
	_jsii_.Get(
		j,
		"reviewerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetRulesPullRequestRequiredReviewersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetRulesPullRequestRequiredReviewersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-github.organizationRuleset.OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOrganizationRulesetRulesPullRequestRequiredReviewersOutputReference_Override(o OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-github.organizationRuleset.OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetFilePatterns(val *[]*string) {
	if err := j.validateSetFilePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePatterns",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetMinimumApprovals(val *float64) {
	if err := j.validateSetMinimumApprovalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumApprovals",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) PutReviewer(value *OrganizationRulesetRulesPullRequestRequiredReviewersReviewer) {
	if err := o.validatePutReviewerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putReviewer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestRequiredReviewersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

