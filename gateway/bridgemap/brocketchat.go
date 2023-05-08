// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/KelvinTegelaar/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
