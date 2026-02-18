// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repositoryenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryEnvironmentReviewersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RepositoryEnvironmentReviewersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RepositoryEnvironmentReviewersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRepositoryEnvironmentReviewersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

