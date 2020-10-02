package v2

import (
	"context"
	"fmt"
)

type CharacterGuild struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

type CharacterData struct {
	Name              string          `json:"name"`
	Title             string          `json:"title"`
	Sex               string          `json:"sex"`
	Vocation          string          `json:"vocation"`
	Level             int             `json:"level"`
	AchievementPoints int             `json:"achievement_points"`
	World             string          `json:"world"`
	Residence         string          `json:"residence"`
	MarriedTo         string          `json:"married_to"`
	Guild             *CharacterGuild `json:"guild"`
	LastLogin         []*Timezone     `json:"last_login"`
	Comment           string          `json:"comment"`
	AccountStatus     string          `json:"account_status"`
	Status            string          `json:"status"`
}

type CharacterAchievements struct {
	Stars int    `json:"stars"`
	Name  string `json:"name"`
}

type CharacterAccountInformation struct {
	LoyaltyTitle string    `json:"loyalty_title"`
	Created      *Timezone `json:"created"`
}

type CharacterOtherCharacters struct {
	Name   string `json:"name"`
	World  string `json:"world"`
	Status string `json:"status"`
}

type Character struct {
	Data         *CharacterData           `json:"data"`
	Achievements []*CharacterAchievements `json:"achievements"`
	//Deaths             []interface{}            `json:"deaths"` //TODO: custom unmarshler
	//AccountInformation *CharacterAccountInformation `json:"account_information"` TODO: custom unmarshler
	//OtherCharacters []*CharacterOtherCharacters `json:"other_characters"` TODO: custom unmarshler
}

type CharacterResponse struct {
	Characters  *Character   `json:"characters"`
	Information *Information `json:"information"`
}

func (c client) Character(context context.Context, name string) (*CharacterResponse, error) {
	characterResponse := &CharacterResponse{}
	url := tibiaDataURL(fmt.Sprintf("characters/%s.json", name))
	err := c.client.Get(context, url, characterResponse)

	return characterResponse, err
}
