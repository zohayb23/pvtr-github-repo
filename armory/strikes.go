package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

// -----
// Strike and Movements for AC_01
// -----

// AC_01 conforms to the Strike function type
func AC_01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "AC_01"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project's version control system MUST require multi-factor authentication for collaborators modifying the project repository settings or accessing sensitive data.",
		ControlID:   "OSPS-AC-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_01_T01)

	return
}

// TODO
func AC_01_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return
}

// -----
// Strike and Movements for AC_01
// -----

// AC_01 conforms to the Strike function type
func AC_02() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "AC_02"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s version control system MUST restrict collaborator permissions to the lowest available privileges by default.",
		ControlID:   "OSPS-AC-01",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(AC_01_T01)

	return
}

// TODO
func AC_01_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to AC_01
	return
}
