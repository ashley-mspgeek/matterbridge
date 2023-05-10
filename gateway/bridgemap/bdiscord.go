// +build !nodiscord

package bridgemap

import (
	bdiscord "github.com/ashley_mspgeek/matterbridge/bridge/discord"
)

func init() {
	FullMap["discord"] = bdiscord.New
	UserTypingSupport["discord"] = struct{}{}
}
