// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/ashley_mspgeek/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
