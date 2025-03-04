package reusable_steps

import (
	"fmt"

	"github.com/revanite-io/sci/pkg/layer4"

	"github.com/revanite-io/pvtr-github-repo/data"
)

func PayloadCheck(payloadData interface{}) (payload data.Payload, message string) {
	payload, ok := payloadData.(data.Payload)
	if !ok {
		message = fmt.Sprintf("Malformed assessment: expected payload type %T, got %T (%v)", data.Payload{}, payloadData, payloadData)
	}
	return
}

func NotImplemented(payloadData interface{}, changes map[string]*layer4.Change) (result layer4.Result, message string) {
	return layer4.Unknown, "Not implemented"
}
