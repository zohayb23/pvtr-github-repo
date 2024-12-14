package armory

import (
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	Armory = raidengine.Armory{
		Tactics: map[string][]raidengine.Strike{
			"maturity_1": {},
			"maturity_2": {},
			"maturity_3": {},
		},
	}
)
