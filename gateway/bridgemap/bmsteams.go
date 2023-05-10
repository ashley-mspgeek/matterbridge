// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/ashley_mspgeek/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
