package tibia

import (
	"fmt"
	"strings"
)

type Location string

const (
	LocationEurope       Location = "Europe"
	LocationSouthAmerica Location = "South America"
	LocationNorthAmerica Location = "North America"
)

func (l Location) String() string {
	return string(l)
}

func (l *Location) Unmarshal(b []byte) error {
	location := Location(strings.Trim(string(b), `"`))
	switch location {
	case LocationEurope, LocationSouthAmerica, LocationNorthAmerica:
		*l = location
		return nil
	default:
		return fmt.Errorf("location doesn't exist: %s", location)
	}
}
