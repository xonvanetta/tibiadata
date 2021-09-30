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

func (wr WorldResponse) validate() error {
	if wr.World.WorldInformation.OnlineRecord == nil {
		return fmt.Errorf("missing online record")
	}
	if wr.World.WorldInformation.PvpType == nil {
		return fmt.Errorf("missing pvp type")
	}
	if wr.World.WorldInformation.Location == nil {
		return fmt.Errorf("missing location")
	}
	return nil
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
	var err error
	worldResponse := &WorldResponse{}
	url := tibiaDataURL(fmt.Sprintf("world/%s.json", name))
	err = c.client.Get(context, url, worldResponse)
	if err != nil {
		return nil, err
	}
	err = worldResponse.validate()
	return worldResponse, err
}
