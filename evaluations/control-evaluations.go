package evaluations

import (
	"github.com/revanite-io/sci/pkg/layer4"
)

//
// Access Control Control Family

func OSPS_AC_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-AC-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-AC-01.01",
		"When a user attempts to access a sensitive resource in the project's version control system, the system MUST require the user to complete a multi-factor authentication process.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_AC_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-AC-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-AC-02.01",
		"When a new collaborator is added, the version control system MUST require manual permission assignment, or restrict the collaborator permissions to the lowest available privileges by default.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_AC_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-AC-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-AC-03.01",
		"When a direct commit is attempted on the project's primary branch, an enforcement mechanism MUST prevent the change from being applied.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-AC-03.02",
		"When an attempt is made to delete the project's primary branch, the version control system MUST treat this as a sensitive activity and require explicit confirmation of intent.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_AC_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-AC-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-AC-04.01",
		"When a CI/CD task is executed with no permissions specified, the project's version control system MUST default to the lowest available permissions for all activities in the pipeline.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-AC-04.02",
		"When a job is assigned permissions in a CI/CD pipeline, the source code or configuration MUST only assign the minimum privileges necessary for the corresponding activity.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

//
// Build and Release Control Family

func OSPS_BR_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-01.01",
		"When a CI/CD pipeline accepts an input parameter, that parameter MUST be sanitized and validated prior to use in the pipeline.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-BR-01.02",
		"When a CI/CD pipeline uses a branch name in its functionality, that name value MUST be sanitized and validated prior to use in the pipeline.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_BR_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-02.01",
		"When an official release is created, that release MUST be assigned a unique version identifier.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-BR-02.01",
		"When an official release is created, all assets within that release MUST be clearly associated with the release identifier or another unique identifier for the asset.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_BR_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-03.01",
		"When the project lists a URI as an official project channel, that URI MUST be exclusively delivered using encrypted channels.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-BR-03.02",
		"When the project lists a URI as an official distribution channel, that URI MUST be exclusively delivered using encrypted channels.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_BR_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-04.01",
		"When an official release is created, that release MUST contain a descriptive log of functional and security modifications.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_BR_05() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-05",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-05.01",
		"When a build and release pipeline ingests dependencies, it MUST use standardized tooling where available.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_BR_06() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-BR-06",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-BR-06.01",
		"When an official release is created, that release MUST be signed or accounted for in a signed manifest including each asset's cryptographic hashes.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

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
			reusable_step_example,
		},
	)

	return
}

func OSPS_DO_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-DO-02",
		Remediation_Guide: "",
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
			reusable_step_example,
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
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-DO-03.02",
		"When the project has made a release, the project documentation MUST contain instructions to verify the expected identity of the person or process authoring the software release.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

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
			reusable_step_example,
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
			reusable_step_example,
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
			reusable_step_example,
		},
	)

	return
}

//
// Governance Control Family

func OSPS_GV_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-GV-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-GV-01.01",
		"While active, the project documentation MUST include a list of project members with access to sensitive resources.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-GV-01.02",
		"While active, the project documentation MUST include descriptions of the roles and responsibilities for members of the project.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_GV_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-GV-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-GV-02.01",
		"While active, the project MUST have one or more mechanisms for public discussions about proposed changes and usage obstacles.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_GV_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-GV-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-GV-03.01",
		"While active, the project documentation MUST include an explanation of the contribution process.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-GV-03.02",
		"While active, the project documentation MUST include a guide for code contributors that includes requirements for acceptable contributions.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_GV_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-GV-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-GV-04.01",
		"While active, the project documentation MUST have a policy that code contributors are reviewed prior to granting escalated permissions to sensitive resources.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

//
// Legal Control Family

func OSPS_LE_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-LE-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-LE-01.01",
		"While active, the version control system MUST require all code contributors to assert that they are legally authorized to make the associated contributions on every commit.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_LE_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-LE-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-LE-02.01",
		"While active, the license for the source code MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-LE-02.02",
		"While active, the license for the released software assets MUST meet the OSI Open Source Definition or the FSF Free Software Definition.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_LE_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-LE-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-LE-03.01",
		"While active, the license for the source code MUST be maintained in the corresponding repository's LICENSE file, COPYING file, or LICENSE/ directory.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-LE-03.02",
		"While active, the license for the released software assets MUST be included in the released source code, or in a LICENSE file, COPYING file, or LICENSE/ directory alongside the corresponding release assets.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

//
// Quality Control Family

func OSPS_QA_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-01.01",
		"While active, the project's source code repository MUST be publicly readable at a static URL.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-QA-01.02",
		"The version control system MUST contain a publicly readable record of all changes made, who made the changes, and when the changes were made.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-02.01",
		"When the package management system supports it, the source code repository MUST contain a dependency list that accounts for the direct language dependencies.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-QA-02.02",
		"When the project has made a release, all compiled released software assets MUST be delivered with a software bill of materials.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-03.01",
		"When a commit is made to the primary branch, any automated status checks for commits MUST pass or be manually bypassed.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-04.01",
		"While active, the project documentation MUST contain a list of any codebases that are considered subprojects or additional repositories.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-QA-04.02",
		"When the project has made a release comprising multiple source code repositories, all subprojects MUST enforce security requirements that are as strict or stricter than the primary codebase.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_05() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-05",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-05.01",
		"While active, the version control system MUST NOT contain generated executable artifacts.",
		[]string{
			"Maturity Level 1",
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_06() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-06",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-06.01",
		"Prior to a commit being accepted, the project's CI/CD pipelines MUST run at least one automated test suite to ensure the changes meet expectations.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-QA-06.02",
		"While active, project's documentation MUST clearly document when and how tests are run.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-QA-06.03",
		"While active, the project's documentation MUST include a policy that all major changes to the software produced by the project should add or update tests of the functionality in an automated test suite.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_QA_07() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-QA-07",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-QA-07.01",
		"When a commit is made to the primary branch, the project's version control system MUST require at least one non-author approval of the changes before merging.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

//
// Security Assessment Control Family

func OSPS_SA_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-SA-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-SA-01.01",
		"When the project has made a release, the project documentation MUST include design documentation demonstrating all actions and actors within the system.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_SA_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-SA-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-SA-02.01",
		"When the project has made a release, the project documentation MUST include descriptions of all external software interfaces of the released software assets.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_SA_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-SA-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-SA-03.01",
		"When the project has made a release, the project MUST perform a security assessment to understand the most likely and impactful potential security problems that could occur within the software.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-SA-03.02",
		"When the project has made a release, the project MUST perform a threat modeling and attack surface analysis to understand and protect against attacks on critical code paths, functions, and interactions within the system.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

//
// Vulnerability Management Control Family

func OSPS_VM_01() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-01",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-01.01",
		"While active, the project documentation MUST include a policy for coordinated vulnerability reporting, with a clear timeframe for response.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_VM_02() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-02",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-02.01",
		"While active, the project documentation MUST contain security contacts.",
		[]string{
			"Maturity Level 1",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_VM_03() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-03",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-03.01",
		"While active, the project documentation MUST provide a means for reporting security vulnerabilities privately to the security contacts within the project.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_VM_04() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-04",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-04.01",
		"While active, the project documentation MUST publicly publish data about discovered vulnerabilities.",
		[]string{
			"Maturity Level 2",
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-VM-04.02",
		"While active, any vulnerabilities in the software components not affecting the project MUST be accounted for in a VEX document, augmenting the vulnerability report with non-exploitability details.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_VM_05() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-05",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-05.01",
		"While active, the project documentation MUST include a policy that defines a threshold for remediation of SCA findings related to vulnerabilities and licenses.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-VM-05.02",
		"While active, the project documentation MUST include a policy to address SCA violations prior to any release.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-VM-05.03",
		"While active, all changes to the project's codebase MUST be automatically evaluated against a documented policy for malicious dependencies and known vulnerabilities in dependencies, then blocked in the event of violations, except when declared and suppressed as non-exploitable.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}

func OSPS_VM_06() (evaluation *layer4.ControlEvaluation) {
	evaluation = &layer4.ControlEvaluation{
		Control_Id:        "OSPS-VM-06",
		Remediation_Guide: "",
	}

	evaluation.AddAssessment(
		"OSPS-VM-06.01",
		"While active, the project documentation MUST include a policy that defines a threshold for remediation of SAST findings.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	evaluation.AddAssessment(
		"OSPS-VM-06.02",
		"While active, all changes to the project's codebase MUST be automatically evaluated against a documented policy for security weaknesses and blocked in the event of violations except when declared and suppressed as non-exploitable.",
		[]string{
			"Maturity Level 3",
		},
		[]layer4.AssessmentStep{
			reusable_step_example,
		},
	)

	return
}
