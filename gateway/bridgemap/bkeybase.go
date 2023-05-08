// +build !nokeybase

package bridgemap

import (
	bkeybase "github.com/KelvinTegelaar/matterbridge/bridge/keybase"
)

func init() {
	FullMap["keybase"] = bkeybase.New
}
