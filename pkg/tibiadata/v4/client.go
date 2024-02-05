package v2

import (
	"context"
	"errors"
	"fmt"

	"github.com/xonvanetta/tibiadata/internal/httpclient"
)

type Client interface {
	//Guild(ctx context.Context, name string) (*GuildResponse, error)
	//Guilds(ctx context.Context, world string) (*GuildsResponse, error)

	World(ctx context.Context, name string) (*WorldResponse, error)
	Worlds(ctx context.Context) (*WorldsResponse, error)

	//Character(ctx context.Context, name string) (*CharacterResponse, error)
	//
	//Highscore(ctx context.Context, world, category string, vocation tibia.Vocation) (*HighscoreResponse, error)
	//
	//News(context context.Context, newsId int) (*NewsResponse, error)
}

type client struct {
	client httpclient.Client
}

var (
	URL         = "https://api.tibiadata.com/v4/"
	ErrNotFound = errors.New("tibiadata: not found")
)

func NewClient() Client {
	return client{
		client: httpclient.New(),
	}
}

func tibiaDataURL(path string) string {
	return fmt.Sprintf("%s%s", URL, path)
}

type Information struct {
	APIDetails APIDetails `json:"api"`
	Timestamp  string     `json:"timestamp"`
	Status     Status     `json:"status"`
}
type APIDetails struct {
	Version int    `json:"version"`
	Release string `json:"release"`
	Commit  string `json:"commit"`
}
type Status struct {
	HTTPCode int    `json:"http_code"`
	Error    int    `json:"error,omitempty"`
	Message  string `json:"message,omitempty"`
}
