package quality

import (
	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
	"github.com/revanite-io/sci/pkg/layer4"
)

func repoIsPublic(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if !data.Repository.IsPrivate {
		return layer4.Passed, "Repository is public"
	}

	return layer4.Failed, "Repository is private"
}

func insightsListsRepositories(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.VerifyPayload(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if len(data.Insights.Project.Repositories) > 0 {
		return layer4.Passed, "Insights contains a list of repositories"
	}

	return layer4.Failed, "Insights does NOT contains a list of repositories"
}
