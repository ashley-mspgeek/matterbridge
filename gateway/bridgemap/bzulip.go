// +build !nozulip

package bridgemap

import (
	bzulip "github.com/ashley_mspgeek/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
