// +build !nowhatsapp
// +build !whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/ashley-mspgeek/matterbridge/bridge/whatsapp"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
