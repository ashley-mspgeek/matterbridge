// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/KelvinTegelaar/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
