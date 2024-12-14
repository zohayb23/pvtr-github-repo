package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
	"github.com/privateerproj/privateer-sdk/utils"
)

//
// !
// !!
// !!!
//
// This file is for reference purposes only
// These are not customized or generated for your use case
// Delete this as soon as you start adding your own changes
//
// !!!
// !!
// !
//

var globalObject interface{}

// Example of a strike that calls an invasive and non-invasive movement.
// Any number or combination of movements can be called
func ExampleStrike01() (strikeName string, result raidengine.StrikeResult) {
	// set default return values
	strikeName = "Example_Strike_01"
	result = raidengine.StrikeResult{
		Description: "The service enforces the use of secure transport protocols for all network communications (e.g., TLS 1.2 or higher).",
		Message:     "Strike has not yet started.",              // This message will be overwritten by subsequent movements
		DocsURL:     "https://maintainer.com/docs/raids/DEV",    // This is an optional link to documentation that will help users better understand the strike
		ControlID:   "CCC.C01",                                  // This is the control ID that the strike is testing against
		Movements:   make(map[string]raidengine.MovementResult), // This map will be populated with the results of each movement
		Passed:      false,                                      // This will be updated to true if a movement passes, and back to false if a movement fails
	}

	result.ExecuteMovement(ExampleMovement0101)

	// if a movement relies on another movement to pass, add this type of condition
	if result.Movements["ExampleMovement0101"].Passed {
		// if a movement could potentially cause harm to the target env, flag it as invasive like this
		result.ExecuteInvasiveMovement(ExampleMovement0102)
	}

	return
}

// ExampleMovement0101 does not apply a change to the system
func ExampleMovement0101() (moveResult raidengine.MovementResult) {
	// Pretend we're making some API call or other logic to determine if the movement is applicable
	customLogicResults := true

	moveResult = raidengine.MovementResult{
		Description: "Making an API call to see if HTTPS is enforced.",
		Function:    utils.CallerPath(0), // This allows interested users to jump directly to the code that is executing this movement
		Passed:      customLogicResults,
	}
	return
}

// ExampleMovement0102 applies an invasive change to the system. Not all changes are invasive, but this one is.
// Use ExecuteInvasiveMovement() to ensure it is run only when the user has opted in to potentially destructive changes.
func ExampleMovement0102() (moveResult raidengine.MovementResult) {
	// The functions here can be defined whereever you like
	// If you have a lot of changes or plan to reuse them, you may want to put them in a separate file
	change1 := raidengine.NewChange(
		"targetName",
		"This change should create a new storage object", // For logging purposes. This will be overwritten by the result of a successful apply function.
		applyChange,
		revertChange,
	)

	// Any intended changes should be applied before returning the movement result
	change1.Apply()

	// A future release may have better object handling for objects returned by the change
	// For now, toss it onto a global variable if you need to access it later
	globalObject = change1.TargetObject

	// If the change is not needed for subsequent movements, revert it now
	// A future release will use this logic to multi-thread the revert process
	// Any changes that are not reverted within the movement will be reverted together at the end of the strike
	change1.Revert()

	// Note that we are not setting Passed to true or false. That will be determined by ExecuteMovement() or ExecuteInvasiveMovement()
	moveResult = raidengine.MovementResult{
		Description: "Making an API call to see if HTTPS is enforced.",
		Function:    utils.CallerPath(0), // This allows interested users to jump directly to the code that is executing this movement
		Changes: map[string]*raidengine.Change{
			"TestChange1": change1,
		},
	}
	return
}

// Mock function to simulate applying a change
func applyChange() (modifiedObject interface{}, err error) {
	// Replace with actual logic
	return nil, nil
}

// Mock function to simulate undoing a change
func revertChange() error {
	// Replace with actual logic
	return nil // Return an error here to simulate failure
}
