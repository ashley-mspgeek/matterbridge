// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/KelvinTegelaar/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
