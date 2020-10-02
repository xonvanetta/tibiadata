package v2

import (
	"context"
	"fmt"
	"sync"
	"testing"

	v2 "github.com/xonvanetta/tibiadata/pkg/tibiadata/v2"

	"github.com/stretchr/testify/assert"
)

func info(t *testing.T, format string, args ...interface{}) {
	t.Log(fmt.Sprintf("Integration Test: "+format, args...))
}

func TestAllGuilds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test, Really long test is not verified yet. JSON unmarshal error guaranteed.")
	}
	v2.URL = "https://api.tibiadata.com/v2/"
	client := v2.NewClient()

	worlds, err := client.Worlds(context.Background())
	assert.NoError(t, err)

	errC := make(chan error)
	var wg sync.WaitGroup

	for _, world := range worlds.Worlds.Allworlds {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			info(t, "running world: %s", name)
			guilds, err := client.Guilds(context.Background(), name)
			if err != nil {
				errC <- err
				return
			}

			for _, guild := range guilds.Guilds.Active {
				info(t, "running world: %s, guild: %s", name, guild.Name)
				_, err := client.Guild(context.Background(), guild.Name)
				if err != nil {
					errC <- err
					return
				}
			}
		}(world.Name)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case err := <-errC:
		t.Fatalf("integration test failed: %s", err)
	case <-done:
	}
}

func TestAllWorlds(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test, Really long test is not verified yet. JSON unmarshal error guaranteed.")
	}
	v2.URL = "https://api.tibiadata.com/v2/"
	client := v2.NewClient()

	worlds, err := client.Worlds(context.Background())
	assert.NoError(t, err)

	for _, world := range worlds.Worlds.Allworlds {
		info(t, "running world: %s", world.Name)
		_, err := client.World(context.Background(), world.Name)
		if err != nil {
			assert.NoError(t, err)
			panic(err)
		}
	}
}

func TestAnticaCharacters(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test, Really long test is not verified yet. JSON unmarshal error guaranteed.")
	}

	v2.URL = "https://api.tibiadata.com/v2/"
	client := v2.NewClient()

	world, err := client.World(context.Background(), "Antica")
	assert.NoError(t, err)

	for _, player := range world.World.PlayersOnline {
		_, err := client.Character(context.Background(), player.Name)
		if err != nil {
			assert.NoError(t, err)
			panic(err)
		}
	}
}
