package access_control

import (
	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
)

func orgRequiresMFA(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	payload, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	required := payload.GraphQL.Organization.RequiresTwoFactorAuthentication

	if required {
		return layer4.Passed, "Two-factor authentication is configured as required by the parent organization"
	}
	return layer4.Failed, "Two-factor authentication is NOT configured as required by the parent organization"
}

func githubBuiltIn(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	_, message = reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	return layer4.Passed, "This control is enforced by GitHub for all projects"
}

func branchProtectionRestrictsPushes(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	payload, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	protectionData := payload.GraphQL.Repository.DefaultBranchRef.BranchProtectionRule

	if protectionData.RestrictsPushes {
		result = layer4.Passed
		message = "Branch protection rule restricts pushes"
	} else if protectionData.RequiresApprovingReviews {
		result = layer4.Passed
		message = "Branch protection rule requires approving reviews"
	} else {
		result = layer4.NeedsReview
		message = "Branch protection rule does not restrict pushes or require approving reviews"
	}
	return
}

func branchProtectionPreventsDeletion(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	payload, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	allowsDeletion := payload.GraphQL.Repository.DefaultBranchRef.RefUpdateRule.AllowsDeletions

	if allowsDeletion {
		result = layer4.Failed
		message = "Branch protection rule allows deletions"
	} else {
		result = layer4.Passed
		message = "Branch protection rule prevents deletions"
	}
	return
}
