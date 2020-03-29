package e2e

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xonvanetta/tibiadata/tibia"
	v2 "github.com/xonvanetta/tibiadata/v2"
)

func mockServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("./tibiadata/" + r.URL.Path)
		assert.NoError(t, err)

		_, err = io.Copy(w, file)
		assert.NoError(t, err)
	}))
}

func TestEndpoints(t *testing.T) {
	tests := []struct {
		name   string
		caller func(t *testing.T, client v2.Client)
	}{
		{
			name: "worlds",
			caller: func(t *testing.T, client v2.Client) {
				response, err := client.Worlds(context.Background())

				assert.NoError(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, 10964, response.Worlds.Online)

				assert.Equal(t, "Antica", response.Worlds.Allworlds[0].Name)
				assert.Equal(t, tibia.LocationEurope, response.Worlds.Allworlds[0].Location)
				assert.Equal(t, tibia.OpenPvP, response.Worlds.Allworlds[0].Worldtype)
				assert.Equal(t, 592, response.Worlds.Allworlds[0].Online)

				assert.Equal(t, "Firmera", response.Worlds.Allworlds[20].Name)
				assert.Equal(t, tibia.LocationNorthAmerica, response.Worlds.Allworlds[20].Location)
				assert.Equal(t, tibia.RetroOpenPvP, response.Worlds.Allworlds[20].Worldtype)
				assert.Equal(t, 89, response.Worlds.Allworlds[20].Online)

				assert.Equal(t, "Noctera", response.Worlds.Allworlds[42].Name)
				assert.Equal(t, tibia.LocationNorthAmerica, response.Worlds.Allworlds[42].Location)
				assert.Equal(t, tibia.OpenPvP, response.Worlds.Allworlds[42].Worldtype)
				assert.Equal(t, 68, response.Worlds.Allworlds[42].Online)
			},
		},
		{
			name: "world/Antica",
			caller: func(t *testing.T, client v2.Client) {
				response, err := client.World(context.Background(), "Antica")

				assert.NoError(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, "Antica", response.World.WorldInformation.Name)
				assert.Equal(t, 589, response.World.WorldInformation.PlayersOnline)
				assert.Equal(t, tibia.LocationEurope, response.World.WorldInformation.Location)
				assert.Equal(t, tibia.OpenPvP, response.World.WorldInformation.PvpType)

				assert.Equal(t, "Don Berco", response.World.PlayersOnline[123].Name)
				assert.Equal(t, tibia.VocationElderDruid, response.World.PlayersOnline[123].Vocation)
				assert.Equal(t, 323, response.World.PlayersOnline[123].Level)

				assert.Equal(t, "Bugga", response.World.PlayersOnline[69].Name)
				assert.Equal(t, tibia.VocationEliteKnight, response.World.PlayersOnline[69].Vocation)
				assert.Equal(t, 194, response.World.PlayersOnline[69].Level)

				assert.Equal(t, "Mini Wziu", response.World.PlayersOnline[349].Name)
				assert.Equal(t, tibia.VocationMasterSorcerer, response.World.PlayersOnline[349].Vocation)
				assert.Equal(t, 27, response.World.PlayersOnline[349].Level)
			},
		},
		{
			name: "world/Secura",
			caller: func(t *testing.T, client v2.Client) {
				response, err := client.World(context.Background(), "Secura")

				assert.NoError(t, err)
				assert.NotNil(t, response)

				assert.Equal(t, "Secura", response.World.WorldInformation.Name)
				assert.Equal(t, 416, response.World.WorldInformation.PlayersOnline)
				assert.Equal(t, tibia.LocationEurope, response.World.WorldInformation.Location)
				assert.Equal(t, tibia.OptionalPvP, response.World.WorldInformation.PvpType)

				assert.Equal(t, "Fenia Tegisa", response.World.PlayersOnline[123].Name)
				assert.Equal(t, tibia.VocationElderDruid, response.World.PlayersOnline[123].Vocation)
				assert.Equal(t, 625, response.World.PlayersOnline[123].Level)

				assert.Equal(t, "Colombo de Vere", response.World.PlayersOnline[69].Name)
				assert.Equal(t, tibia.VocationEliteKnight, response.World.PlayersOnline[69].Vocation)
				assert.Equal(t, 941, response.World.PlayersOnline[69].Level)

				assert.Equal(t, "Sylikerago Yekerlion", response.World.PlayersOnline[349].Name)
				assert.Equal(t, tibia.VocationRoyalPaladin, response.World.PlayersOnline[349].Vocation)
				assert.Equal(t, 385, response.World.PlayersOnline[349].Level)
			},
		},
		{
			name: "guilds/Antica",
			caller: func(t *testing.T, client v2.Client) {
				response, err := client.Guilds(context.Background(), "Antica")

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "Antica", response.Guilds.World)

				assert.Equal(t, "Circle of Mages", response.Guilds.Active[20].Name)
				assert.Equal(t, "Hello Wanderer", response.Guilds.Active[20].Desc)

				assert.Equal(t, "Green Dragon", response.Guilds.Active[40].Name)
				assert.Equal(t, "The Green Dragon Inn is located directly north of Thais depot. Services include a bar with all rights, a casino with sporadic opening hours, rentable rooms and beds. We're always looking for new hires (Min. lvl: 20) to help run the inn! Visit our discord channel after you join!", response.Guilds.Active[40].Desc)

				assert.Equal(t, "Neck Stab", response.Guilds.Active[56].Name)
				assert.Equal(t, "", response.Guilds.Active[56].Desc)
			},
		},
		{
			name: "guild/Red Rose",
			caller: func(t *testing.T, client v2.Client) {
				response, err := client.Guild(context.Background(), "Red Rose")

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "Antica", response.Guild.Data.World)

				assert.Equal(t, "Red Rose", response.Guild.Data.Name)
				assert.Equal(t, 121, response.Guild.Data.Totalmembers)
				assert.Equal(t, 3, response.Guild.Data.OnlineStatus)
				assert.Equal(t, 118, response.Guild.Data.OfflineStatus)

				assert.Equal(t, "Leader", response.Guild.Members[0].RankTitle)
				assert.Equal(t, "Avora Skyfallen", response.Guild.Members[0].Characters[0].Name)
				assert.Equal(t, tibia.VocationMasterSorcerer, response.Guild.Members[0].Characters[0].Vocation)
				assert.Equal(t, 239, response.Guild.Members[0].Characters[0].Level)
				assert.Equal(t, "offline", response.Guild.Members[0].Characters[0].Status)

				assert.Equal(t, "Praetor", response.Guild.Members[1].RankTitle)
				assert.Equal(t, "Miujau", response.Guild.Members[1].Characters[0].Name)
				assert.Equal(t, tibia.VocationEliteKnight, response.Guild.Members[1].Characters[0].Vocation)
				assert.Equal(t, 321, response.Guild.Members[1].Characters[0].Level)
				assert.Equal(t, "offline", response.Guild.Members[1].Characters[0].Status)
			},
		},
	}

	server := mockServer(t)
	defer server.Close()
	v2.URL = server.URL + "/"

	client := v2.New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.caller(t, client)
		})
	}
}
