// +build !nozulip

package bridgemap

import (
	bzulip "github.com/KelvinTegelaar/matterbridge/bridge/zulip"
)

func init() {
	FullMap["zulip"] = bzulip.New
}
