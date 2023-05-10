// +build !nomumble

package bridgemap

import (
	bmumble "github.com/ashley_mspgeek/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
