// +build !notelegram

package bridgemap

import (
	btelegram "github.com/ashley_mspgeek/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
