// +build !nosteam

package bridgemap

import (
	bsteam "github.com/ashley-mspgeek/matterbridge/bridge/steam"
)

func init() {
	FullMap["steam"] = bsteam.New
}
