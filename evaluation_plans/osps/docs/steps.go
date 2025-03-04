package docs

import (
	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
)

func hasUserGuidees(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Project.Documentation.DetailedGuide == "" {
		return layer4.Failed, "User guide was NOT specified in Security Insights data"
	}

	return layer4.Passed, "User guide was specified in Security Insights data"
}

func hasDefectReportingGuide(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	// data, message := reusable_steps.PayloadCheck(payloadData)
	// if message != "" {
	// 	return layer4.Unknown, message
	// }

	// if data.Insights.Project.Documentation.BugReportGuide == "" {
	// 	return layer4.Failed, "Defect report guide was NOT specified in Security Insights data"
	// }

	return layer4.Unknown, "Defect report guide was specified in Security Insights data"
}

func hasIssuesEnabled(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Repository.HasIssuesEnabled {
		return layer4.Passed, "Issues are enabled for the repository"
	}

	return layer4.Failed, "Issues are disabled for the repository"
}

func hasDiscussionsEnabled(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Repository.HasDiscussionsEnabled {
		return layer4.Passed, "Discussions are enabled for the repository"
	}

	return layer4.Failed, "Discussions are disabled for the repository"
}

func acceptsVulnReports(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Project.Vulnerability.ReportsAccepted {
		return layer4.Passed, "Repository accepts vulnerability reports"
	}

	return layer4.Failed, "Repository does not accept vulnerability reports"
}

func hasSignatureVerificationGuide(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
	data, message := reusable_steps.PayloadCheck(payloadData)
	if message != "" {
		return layer4.Unknown, message
	}

	if data.Insights.Project.Documentation.SignatureVerification == "" {
		return layer4.Failed, "Signature verification guide was NOT specified in Security Insights data"
	}

	return layer4.Passed, "Signature verification guide was specified in Security Insights data"
}
