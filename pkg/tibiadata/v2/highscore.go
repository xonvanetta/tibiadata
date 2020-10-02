package v2

import (
	"context"
	"fmt"

	"github.com/xonvanetta/tibiadata/pkg/tibia"
)

type HighscoreFilters struct {
	World    string `json:"world"`
	Category string `json:"category"`
	Vocation string `json:"vocation"`
}

type HighscoreData struct {
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	World    string `json:"world"`
	Vocation string `json:"vocation"`
	Points   int64  `json:"points"`
	Level    int    `json:"level"`
}

type Highscore struct {
	Filters *HighscoreFilters `json:"filters"`
	Data    []*HighscoreData  `json:"data"`
}

type HighscoreResponse struct {
	Highscores  *Highscore   `json:"highscores"`
	Information *Information `json:"information"`
}

func (c client) Highscore(context context.Context, world, category string, vocation tibia.Vocation) (*HighscoreResponse, error) {
	highscoreResponse := &HighscoreResponse{}

	url := fmt.Sprintf("highscores/%s", world)
	if category != "" {
		url = fmt.Sprintf("%s/%s", url, category)
	}
	if vocation != "" {
		url = fmt.Sprintf("%s/%s", url, vocation.String())
	}

	url = tibiaDataURL(fmt.Sprintf("%s.json", url))
	err := c.client.Get(context, url, highscoreResponse)

	return highscoreResponse, err
}
