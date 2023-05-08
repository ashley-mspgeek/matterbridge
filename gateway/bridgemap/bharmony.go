//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/KelvinTegelaar/matterbridge/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}
