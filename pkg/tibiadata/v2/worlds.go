package v2

import (
	"context"

	"github.com/xonvanetta/tibiadata/pkg/tibia"
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

func (c client) Worlds(context context.Context) (WorldsResponse, error) {
	var worldsResponse WorldsResponse
	err := c.client.Get(context, tibiaDataURL("worlds.json"), &worldsResponse)
	return worldsResponse, err
}
