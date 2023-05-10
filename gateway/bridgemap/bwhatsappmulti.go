// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/ashley_mspgeek/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
