// +build !noapi

package bridgemap

import (
	"github.com/ashley-mspgeek/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
