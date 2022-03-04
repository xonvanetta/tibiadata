package v2

import (
	"context"
	"time"
)

type WorldsResponse struct {
	Worlds *struct {
		PlayersOnline int       `json:"players_online"`
		RecordPlayers int       `json:"record_players"`
		RecordDate    time.Time `json:"record_date"`
		RegularWorlds []*struct {
			Name                string `json:"name"`
			Status              string `json:"status"`
			PlayersOnline       int    `json:"players_online"`
			Location            string `json:"location"`
			PvpType             string `json:"pvp_type"`
			PremiumOnly         bool   `json:"premium_only"`
			TransferType        string `json:"transfer_type"`
			BattleyeProtected   bool   `json:"battleye_protected"`
			BattleyeDate        string `json:"battleye_date"`
			GameWorldType       string `json:"game_world_type"`
			TournamentWorldType string `json:"tournament_world_type"`
		} `json:"regular_worlds"`
		TournamentWorlds []*struct {
			Name                string `json:"name"`
			Status              string `json:"status"`
			PlayersOnline       int    `json:"players_online"`
			Location            string `json:"location"`
			PvpType             string `json:"pvp_type"`
			PremiumOnly         bool   `json:"premium_only"`
			TransferType        string `json:"transfer_type"`
			BattleyeProtected   bool   `json:"battleye_protected"`
			BattleyeDate        string `json:"battleye_date"`
			GameWorldType       string `json:"game_world_type"`
			TournamentWorldType string `json:"tournament_world_type"`
		} `json:"tournament_worlds"`
	} `json:"worlds"`
	Information *Information `json:"information"`
}

func (c client) Worlds(context context.Context) (*WorldsResponse, error) {
	worldsResponse := &WorldsResponse{}
	err := c.client.Get(context, tibiaDataURL("worlds"), worldsResponse)
	return worldsResponse, err
}
