// +build !noxmpp

package bridgemap

import (
	bxmpp "github.com/ashley-mspgeek/matterbridge/bridge/xmpp"
)

func init() {
	FullMap["xmpp"] = bxmpp.New
}
