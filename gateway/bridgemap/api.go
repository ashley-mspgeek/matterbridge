// +build !noapi

package bridgemap

import (
	"github.com/KelvinTegelaar/matterbridge/bridge/api"
)

func init() {
	FullMap["api"] = api.New
}
