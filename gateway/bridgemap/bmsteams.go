// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/KelvinTegelaar/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
