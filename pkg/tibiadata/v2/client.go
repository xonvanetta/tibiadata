package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/xonvanetta/tibiadata/internal/httpclient"
)

type Client interface {
	Guild(ctx context.Context, name string) (*GuildResponse, error)
	Guilds(ctx context.Context, world string) (*GuildsResponse, error)

	World(ctx context.Context, name string) (*WorldResponse, error)
	Worlds(ctx context.Context) (*WorldsResponse, error)

	Character(ctx context.Context, name string) (*CharacterResponse, error)
}

type client struct {
	client httpclient.Client
}

var (
	location *time.Location

	URL         = "https://api.tibiadata.com/v2/"
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
	APIVersion    int     `json:"api_version"`
	ExecutionTime float64 `json:"execution_time"`
	LastUpdated   Time    `json:"last_updated"`
	Timestamp     Time    `json:"timestamp"`
}

type Timezone struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(b []byte) error {
	var err error
	if location == nil {
		location, err = time.LoadLocation("Europe/Stockholm")
		if err != nil {
			panic(fmt.Errorf("tibiadata: failed to load location: %w", err))
		}
	}

	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), location)
	return err
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02"`, string(b))
	return err
}
