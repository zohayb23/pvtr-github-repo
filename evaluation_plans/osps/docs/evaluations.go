package docs // not sure why, but 'documentation' was misbehaving as a package name. reserved?

import (
	"github.com/revanite-io/pvtr-github-repo/evaluation_plans/reusable_steps"
	"github.com/revanite-io/sci/pkg/layer4"
)

//
// Documentation Control Family

func OSPS_DO_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-DO-01.01",
		"When the project has made a release, the project documentation MUST include user guides for all basic functionality.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.HasMadeReleases,
			hasUserGuidees,
		},
	)

	return
}

func OSPS_DO_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-02",
		Remediation_Guide: "",
	}

	issuesOrDiscussionsEnabled := func(payloadData interface{}, _ map[string]*layer4.Change) (result layer4.Result, message string) {
		result, message = hasIssuesEnabled(payloadData, nil)
		if result == layer4.Passed {
			return
		}
		result, message = hasDiscussionsEnabled(payloadData, nil)
		if result == layer4.Passed {
			return
		}
		return layer4.Failed, "Both issues and discussions are disabled for the repository"
	}

	evaluation.AddAssessment(
		"OSPS-DO-02.01",
		"When the project has made a release, the project documentation MUST include a guide for reporting defects.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.HasMadeReleases,
			issuesOrDiscussionsEnabled,
			acceptsVulnReports,
		},
	)

	return
}

func OSPS_DO_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-DO-03.01",
		"When the project has made a release, the project documentation MUST contain instructions to verify the integrity and authenticity of the release assets.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.HasMadeReleases,
			hasSignatureVerificationGuide,
		},
	)

	// TODO: Implement this assessment
	// evaluation.AddAssessment(
	// 	"OSPS-DO-03.02",
	// 	"When the project has made a release, the project documentation MUST contain instructions to verify the expected identity of the person or process authoring the software release.",
	// 	[]string{
	// 		"Maturity Level 3",
	// 	},
	// 	[]layer4.AssessmentStep{
	// 		reusable_steps.NotImplemented,
	// 	},
	// )

	return
}

func OSPS_DO_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-DO-04.01",
		"When the project has made a release, the project documentation MUST include a descriptive statement about the scope and duration of support for each release.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.NotImplemented,
		},
	)

	return
}

func OSPS_DO_05() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-05",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-DO-05.01",
		"When the project has made a release, the project documentation MUST provide a descriptive statement when releases or versions will no longer receive security updates.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.NotImplemented,
		},
	)

	return
}

func OSPS_DO_06() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-06",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-DO-06.01",
		"When the project has made a release, the project documentation MUST include a description of how the project selects, obtains, and tracks its dependencies.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_steps.HasMadeReleases,
			hasDependencyManagementPolicy,
		},
	)

	return
}
