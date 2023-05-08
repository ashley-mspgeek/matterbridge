// +build !novk

package bridgemap

import (
	bvk "github.com/KelvinTegelaar/matterbridge/bridge/vk"
)

func init() {
	FullMap["vk"] = bvk.New
}
