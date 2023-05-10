// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/ashley-mspgeek/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
