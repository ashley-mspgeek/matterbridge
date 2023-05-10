// +build !nozulip

package bridgemap

import (
	bzulip "github.com/ashley-mspgeek/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
