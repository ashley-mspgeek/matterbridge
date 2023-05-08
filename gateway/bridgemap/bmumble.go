// +build !nomumble

package bridgemap

import (
	bmumble "github.com/KelvinTegelaar/matterbridge/bridge/mumble"
)

func init() {
	FullMap["mumble"] = bmumble.New
}
