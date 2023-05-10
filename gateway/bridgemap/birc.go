// +build !noirc

package bridgemap

import (
	birc "github.com/ashley_mspgeek/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
