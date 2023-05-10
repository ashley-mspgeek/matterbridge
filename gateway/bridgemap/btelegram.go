// +build !notelegram

package bridgemap

import (
	btelegram "github.com/ashley-mspgeek/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
