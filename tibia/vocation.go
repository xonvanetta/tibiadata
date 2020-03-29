package tibia

import (
	"fmt"
	"strings"
)

type Vocation string

const (
	VocationNone           Vocation = "None"
	VocationDruid          Vocation = "Druid"
	VocationKnight         Vocation = "Knight"
	VocationPaladin        Vocation = "Paladin"
	VocationSorcerer       Vocation = "Sorcerer"
	VocationElderDruid     Vocation = "Elder Druid"
	VocationEliteKnight    Vocation = "Elite Knight"
	VocationRoyalPaladin   Vocation = "Royal Paladin"
	VocationMasterSorcerer Vocation = "Master Sorcerer"
)

func (v *Vocation) Unmarshal(b []byte) error {
	vocation := Vocation(strings.Trim(string(b), `"`))
	switch vocation {
	case VocationNone, VocationDruid, VocationKnight, VocationPaladin, VocationSorcerer, VocationElderDruid, VocationEliteKnight, VocationRoyalPaladin, VocationMasterSorcerer:
		*v = vocation
		return nil
	default:
		return fmt.Errorf("vocation doesn't exist: %s", vocation)
	}
}
