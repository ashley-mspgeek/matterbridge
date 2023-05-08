// +build !nomatrix

package bridgemap

import (
	bmatrix "github.com/KelvinTegelaar/matterbridge/bridge/matrix"
)

func init() {
	FullMap["matrix"] = bmatrix.New
}
