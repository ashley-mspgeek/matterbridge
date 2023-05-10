// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/ashley_mspgeek/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
