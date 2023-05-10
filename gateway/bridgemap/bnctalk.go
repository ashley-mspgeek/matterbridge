// +build !nonctalk

package bridgemap

import (
	btalk "github.com/ashley_mspgeek/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
