// +build !norocketchat

package bridgemap

import (
	brocketchat "github.com/ashley-mspgeek/matterbridge/bridge/rocketchat"
)

func init() {
	FullMap["rocketchat"] = brocketchat.New
}
