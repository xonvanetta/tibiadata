package v2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/xonvanetta/tibiadata/tibia"
)

type Guild struct {
	Error   string         `json:"error"`
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
	Invited Date   `json:"invited"`
}

type GuildCharacter struct {
	Name     string         `json:"name"`
	Nick     string         `json:"nick"`
	Level    uint64         `json:"level"`
	Vocation tibia.Vocation `json:"vocation"`
	Joined   Date           `json:"joined"`
	Online   Online         `json:"status"`
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02"`, string(b))
	return err
}

type Online bool

func (o Online) Bool() bool {
	return bool(o)
}

func (o *Online) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"online"`:
		*o = true
	case `"offline"`:
		*o = false
	default:
		return fmt.Errorf("online status doesnt exist: %s", string(b))
	}

	return nil
}

type GuildData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	//Guildhall     GuildHall `json:"guildhall"`
	Application   bool `json:"application"`
	War           bool `json:"war"`
	OnlineStatus  int  `json:"online_status"`
	OfflineStatus int  `json:"offline_status"`
	//Disbanded     bool      `json:"disbanded"`
	Totalmembers int    `json:"totalmembers"`
	Totalinvited int    `json:"totalinvited"`
	World        string `json:"world"`
	Founded      string `json:"founded"`
	Active       bool   `json:"active"`
	Guildlogo    string `json:"guildlogo"`
}

//type GuildHall struct {
//	Name    string `json:"name"`
//	Town    string `json:"town"`
//	Paid    string `json:"paid"`
//	World   string `json:"world"`
//	Houseid int    `json:"houseid"`
//}

type GuildResponse struct {
	Guild       Guild       `json:"guild"`
	Information Information `json:"information"`
}

var (
	ErrNotFound = errors.New("tibiadata: not found")
)

func (c Client) Guild(context context.Context, name string) (GuildResponse, error) {
	var guildResponse GuildResponse
	url := fmt.Sprintf("guild/%s.json", name)
	err := c.get(context, url, &guildResponse)
	if guildResponse.Guild.Error != "" {
		return guildResponse, guildToError(guildResponse.Guild.Error)
	}

	return guildResponse, err
}

func guildToError(err string) error {
	switch err {
	case "Guild does not exist.":
		return ErrNotFound
	default:
		log.Println(fmt.Sprintf("err not found in error list: %s", err))
		return fmt.Errorf(err)
	}
}
