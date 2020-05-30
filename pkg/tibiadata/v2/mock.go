package v2

import "context"

type Mock struct {
	err            error
	guildResponse  *GuildResponse
	guildsResponse *GuildsResponse
	worldResponse  *WorldResponse
	worldsResponse *WorldsResponse
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Guild(context.Context, string) (*GuildResponse, error) {
	return m.guildResponse, m.err
}

func (m *Mock) Guilds(context.Context, string) (*GuildsResponse, error) {
	return m.guildsResponse, m.err
}

func (m *Mock) World(context.Context, string) (*WorldResponse, error) {
	return m.worldResponse, m.err
}

func (m *Mock) Worlds(context.Context) (*WorldsResponse, error) {
	return m.worldsResponse, m.err
}

func (m *Mock) SetGuildResponse(guildResponse *GuildResponse) {
	m.guildResponse = guildResponse
}

func (m *Mock) SetGuildsResponse(guildsResponse *GuildsResponse) {
	m.guildsResponse = guildsResponse
}

func (m *Mock) SetWorldResponse(worldResponse *WorldResponse) {
	m.worldResponse = worldResponse
}

func (m *Mock) SetWorldsResponse(worldsResponse *WorldsResponse) {
	m.worldsResponse = worldsResponse
}

func (m *Mock) SetError(err error) {
	m.err = err
}
