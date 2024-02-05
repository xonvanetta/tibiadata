package v2

import (
	"context"
	"fmt"
	"time"
)

type WorldResponse struct {
	World *struct {
		Name                string    `json:"name"`
		Status              string    `json:"status"`
		PlayersOnline       int       `json:"players_online"`
		RecordPlayers       int       `json:"record_players"`
		RecordDate          time.Time `json:"record_date"`
		CreationDate        string    `json:"creation_date"`
		Location            string    `json:"location"`
		PvpType             string    `json:"pvp_type"`
		PremiumOnly         bool      `json:"premium_only"`
		TransferType        string    `json:"transfer_type"`
		WorldQuestTitles    []string  `json:"world_quest_titles"`
		BattleyeProtected   bool      `json:"battleye_protected"`
		BattleyeDate        string    `json:"battleye_date"`
		GameWorldType       string    `json:"game_world_type"`
		TournamentWorldType string    `json:"tournament_world_type"`
		OnlinePlayers       []*struct {
			Name     string `json:"name"`
			Level    int    `json:"level"`
			Vocation string `json:"vocation"`
		} `json:"online_players"`
	} `json:"world"`
	Information *Information `json:"information"`
}

func (c client) World(context context.Context, name string) (*WorldResponse, error) {
	worldResponse := &WorldResponse{}
	url := tibiaDataURL(fmt.Sprintf("world/%s", name))
	err := c.client.Get(context, url, worldResponse)
	return worldResponse, err
}
