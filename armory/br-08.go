package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

func BR_08() (strikeName string, result raidengine.StrikeResult) {
	strikeName = "BR_08"
	result = raidengine.StrikeResult{
		Passed:      false,
		Description: "The project’s version control system MUST restrict collaborator permissions to the lowest available privileges by default.",
		ControlID:   "OSPS-BR-08",
		Movements:   make(map[string]raidengine.MovementResult),
	}

	result.ExecuteMovement(BR_08_T01)

	return
}

func BR_08_T01() (moveResult raidengine.MovementResult) {
	moveResult = raidengine.MovementResult{
		Description: "This movement is still under construction",
		Function:    utils.CallerPath(0),
	}

	// TODO: Use this section to write a single step or test that contributes to BR_01
	return
}
