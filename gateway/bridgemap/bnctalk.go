// +build !nonctalk

package bridgemap

import (
	btalk "github.com/ashley-mspgeek/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
