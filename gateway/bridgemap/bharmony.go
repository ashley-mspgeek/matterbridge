//go:build !noharmony
// +build !noharmony

package bridgemap

import (
	bharmony "github.com/ashley_mspgeek/matterbridge/bridge/harmony"
)

func init() {
	FullMap["harmony"] = bharmony.New
}
