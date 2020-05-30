package v2

import (
	"context"
	"fmt"
)

type Guilds struct {
	World     string             `json:"world"`
	Active    []GuildInformation `json:"active"`
	Formation []GuildInformation `json:"formation"`
}

type GuildInformation struct {
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Guildlogo string `json:"guildlogo"`
}

type GuildsResponse struct {
	Guilds      Guilds      `json:"guilds"`
	Information Information `json:"information"`
}

func (c client) Guilds(context context.Context, world string) (*GuildsResponse, error) {
	guildsResponse := &GuildsResponse{}
	url := tibiaDataURL(fmt.Sprintf("guilds/%s.json", world))
	err := c.client.Get(context, url, guildsResponse)
	return guildsResponse, err
}
