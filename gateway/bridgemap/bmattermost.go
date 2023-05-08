// +build !nomattermost

package bridgemap

import (
	bmattermost "github.com/KelvinTegelaar/matterbridge/bridge/mattermost"
)

func init() {
	FullMap["mattermost"] = bmattermost.New
}
