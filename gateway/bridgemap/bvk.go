// +build !novk

package bridgemap

import (
	bvk "github.com/ashley-mspgeek/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
