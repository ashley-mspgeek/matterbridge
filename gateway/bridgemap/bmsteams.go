// +build !nomsteams

package bridgemap

import (
	bmsteams "github.com/ashley-mspgeek/matterbridge/bridge/msteams"
)

func init() {
	FullMap["msteams"] = bmsteams.New
}
