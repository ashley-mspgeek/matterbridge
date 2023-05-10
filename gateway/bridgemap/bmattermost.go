// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/ashley-mspgeek/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
