// +build !nomumble

package bridgemap

import (
	bmumble "github.com/ashley-mspgeek/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
