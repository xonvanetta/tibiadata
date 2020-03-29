package v2

import (
	"context"
	"fmt"

	"github.com/xonvanetta/tibiadata/tibia"
)

type Guild struct {
	Data    GuildData      `json:"data"`
	Members []GuildMembers `json:"members"`
	Invited []GuildInvite  `json:"invited"`
}

type GuildMembers struct {
	RankTitle  string           `json:"rank_title"`
	Characters []GuildCharacter `json:"characters"`
}

type GuildInvite struct {
	Name    string `json:"name"`
	Invited string `json:"invited"` //Todo: Time format 2006-01-02
}

type GuildCharacter struct {
	Name     string         `json:"name"`
	Nick     string         `json:"nick"`
	Level    int            `json:"level"`
	Vocation tibia.Vocation `json:"vocation"`
	Joined   string         `json:"joined"`
	Status   string         `json:"status"` // Todo:  online: bool
}

type GuildData struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Guildhall     GuildHall `json:"guildhall"`
	Application   bool      `json:"application"`
	War           bool      `json:"war"`
	OnlineStatus  int       `json:"online_status"`
	OfflineStatus int       `json:"offline_status"`
	Disbanded     bool      `json:"disbanded"`
	Totalmembers  int       `json:"totalmembers"`
	Totalinvited  int       `json:"totalinvited"`
	World         string    `json:"world"`
	Founded       string    `json:"founded"`
	Active        bool      `json:"active"`
	Guildlogo     string    `json:"guildlogo"`
}

type GuildHall struct {
	Name    string `json:"name"`
	Town    string `json:"town"`
	Paid    string `json:"paid"`
	World   string `json:"world"`
	Houseid int    `json:"houseid"`
}

type GuildResponse struct {
	Guild       Guild       `json:"guild"`
	Information Information `json:"information"`
}

func (c Client) Guild(context context.Context, name string) (GuildResponse, error) {
	var guildResponse GuildResponse
	url := fmt.Sprintf("guild/%s.json", name)
	err := c.get(context, url, &guildResponse)
	return guildResponse, err
}
