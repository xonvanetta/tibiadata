package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

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
	URL         = "https://api.tibiadata.com/v3/"
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
	APIVersion int       `json:"api_version"`
	Timestamp  time.Time `json:"timestamp"`
}
