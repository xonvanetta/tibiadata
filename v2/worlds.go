package v2

import (
	"context"

	"github.com/xonvanetta/tibiadata/tibia"
)

type Worlds struct {
	Online    int        `json:"online"`
	Allworlds []AllWorld `json:"allworlds"`
}

type AllWorld struct {
	Name       string         `json:"name"`
	Online     int            `json:"online"`
	Location   tibia.Location `json:"location"`
	Worldtype  tibia.PvPType  `json:"worldtype"`
	Additional string         `json:"additional"`
}

type WorldsResponse struct {
	Worlds      Worlds      `json:"worlds"`
	Information Information `json:"information"`
}

func (c Client) Worlds(context context.Context) (WorldsResponse, error) {
	var worldsResponse WorldsResponse
	err := c.get(context, "worlds.json", &worldsResponse)
	return worldsResponse, err
}
