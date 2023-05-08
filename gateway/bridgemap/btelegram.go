// +build !notelegram

package bridgemap

import (
	btelegram "github.com/KelvinTegelaar/matterbridge/bridge/telegram"
)

func init() {
	FullMap["telegram"] = btelegram.New
}
