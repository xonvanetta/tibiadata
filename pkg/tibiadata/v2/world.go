package v2

import (
	"context"
	"fmt"

	"github.com/xonvanetta/tibiadata/pkg/tibia"
)

type WorldResponse struct {
	World       *World       `json:"world"`
	Information *Information `json:"information"`
}

type World struct {
	WorldInformation *WorldInformation `json:"world_information"`
	PlayersOnline    []*PlayerOnline   `json:"players_online"`
}

type WorldInformation struct {
	Name             string          `json:"name"`
	PlayersOnline    int             `json:"players_online"`
	OnlineRecord     *OnlineRecord   `json:"online_record"`
	CreationDate     string          `json:"creation_date"` //1997-01
	Location         *tibia.Location `json:"location"`
	PvpType          *tibia.PvPType  `json:"pvp_type"`
	WorldQuestTitles []string        `json:"world_quest_titles"`
	BattleyeStatus   string          `json:"battleye_status"`
}

type OnlineRecord struct {
	Players int       `json:"players"`
	Date    *Timezone `json:"date"`
}

type PlayerOnline struct {
	Name     string          `json:"name"`
	Level    int             `json:"level"`
	Vocation *tibia.Vocation `json:"vocation"`
}

func (c client) World(context context.Context, name string) (*WorldResponse, error) {
	worldResponse := &WorldResponse{}
	url := tibiaDataURL(fmt.Sprintf("world/%s.json", name))
	err := c.client.Get(context, url, worldResponse)
	return worldResponse, err
}
