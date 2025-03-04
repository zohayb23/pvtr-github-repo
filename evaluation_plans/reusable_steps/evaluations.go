package reusable_steps

import (
	"fmt"

	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/data"
)

// TODO: This is only for reference, and should be deleted
func StepExample(payloadData interface{}, changes map[string]*layer4.Change) (result layer4.Result, message string) {
	payload, ok := payloadData.(data.Payload) // TODO: return the data, not all of it
	if !ok {
		return layer4.Unknown, fmt.Sprintf("Malformed assessment: expected payload type %T, got %T (%v)", data.Payload{}, payloadData, payloadData)
	}
	if payload.GraphQL.Repository.Name != "" {
		return layer4.Passed, fmt.Sprint("Repo Name: ", payload.GraphQL.Repository.Name)
	}
	return layer4.Unknown, "Not implemented"
}
