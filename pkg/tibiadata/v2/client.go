package v2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/xonvanetta/tibiadata/internal/httpclient"
)

type Client interface {
	Guild(context.Context, string) (*GuildResponse, error)
	Guilds(context.Context, string) (*GuildsResponse, error)

	World(context.Context, string) (*WorldResponse, error)
	Worlds(context.Context) (*WorldsResponse, error)
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
