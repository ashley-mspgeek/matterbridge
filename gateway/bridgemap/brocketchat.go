// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/ashley_mspgeek/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
