// +build !noapi

package bridgemap

import (
	"github.com/ashley_mspgeek/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
