package tibia

import (
	"fmt"
	"strings"
)

type PvPType string

const (
	OpenPvP          PvPType = "Open PvP"
	OptionalPvP      PvPType = "Optional PvP"
	HardcorePvP      PvPType = "Hardcore PvP"
	RetroHardcorePvP PvPType = "Retro Hardcore PvP"
	RetroOpenPvP     PvPType = "Retro Open PvP"
)

func (t PvPType) String() string {
	return string(t)
}

func (t *PvPType) UnmarshalJSON(b []byte) error {
	worldType := PvPType(strings.Trim(string(b), `"`))
	switch worldType {
	case OpenPvP, OptionalPvP, HardcorePvP, RetroHardcorePvP, RetroOpenPvP:
		*t = worldType
		return nil
	default:
		return fmt.Errorf("worldtype doesn't exist: %s", worldType)
	}
}
