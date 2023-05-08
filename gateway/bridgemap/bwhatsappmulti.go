// +build whatsappmulti

package bridgemap

import (
	bwhatsapp "github.com/KelvinTegelaar/matterbridge/bridge/whatsappmulti"
)

func init() {
	FullMap["whatsapp"] = bwhatsapp.New
}
