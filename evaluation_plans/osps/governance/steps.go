package governance

import (
	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
	"github.com/revanite-io/sci/pkg/layer4"
)

func coreTeamIsListed(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if len(data.Insights.Repository.CoreTeam) == 0 {
		return layer4.Failed, "Core team was NOT specified in Security Insights data"
	}

	return layer4.Passed, "Core team was specified in Security Insights data"
}

func projectAdminsListed(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if len(data.Insights.Project.Administrators) == 0 {
		return layer4.Failed, "Project admins were NOT specified in Security Insights data"
	}

	return layer4.Passed, "Project admins were specified in Security Insights data"
}

func hasRolesAndResponsibilities(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Repository.Documentation.Governance == "" {
		return layer4.Failed, "Roles and responsibilities were NOT specified in Security Insights data"
	}

	return layer4.Passed, "Roles and responsibilities were specified in Security Insights data"
}

func hasContributionGuide(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Project.Documentation.CodeOfConduct != "" && data.Insights.Repository.Documentation.Contributing != "" {
		return layer4.Failed, "Contributing guide and code of conduct locations specified in Security Insights data"
	}

	if data.Repository.ContributingGuidelines.Body != "" && data.Insights.Project.Documentation.CodeOfConduct != "" {
		return layer4.Passed, "Contributing guide was found via GitHub API and code of conduct was specified in Security Insights data"
	}

	if data.Repository.ContributingGuidelines.Body != "" {
		return layer4.NeedsReview, "Contributing guide was found via GitHub API, but code of conduct was NOT specified in Security Insights data"
	}

	return layer4.Failed, "Contribution guide not found in Security Insights data or via GitHub API"
}
