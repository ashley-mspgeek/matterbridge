// +build !noirc

package bridgemap

import (
	birc "github.com/KelvinTegelaar/matterbridge/bridge/irc"
)

func init() {
	FullMap["irc"] = birc.New
}
