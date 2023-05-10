// +build !noirc

package bridgemap

import (
	birc "github.com/ashley-mspgeek/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
