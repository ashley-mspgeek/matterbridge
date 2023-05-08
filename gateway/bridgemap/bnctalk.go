// +build !nonctalk

package bridgemap

import (
	btalk "github.com/KelvinTegelaar/matterbridge/bridge/nctalk"
)

func init() {
	FullMap["nctalk"] = btalk.New
}
