// +build !nosshchat

package bridgemap

import (
	bsshchat "github.com/ashley-mspgeek/matterbridge/bridge/sshchat"
)

func init() {
	FullMap["sshchat"] = bsshchat.New
}
