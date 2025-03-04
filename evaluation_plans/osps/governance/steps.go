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
