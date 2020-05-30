package v2

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	v2 "github.com/xonvanetta/tibiadata/pkg/tibiadata/v2"
)

var name string

func BenchmarkGuild(b *testing.B) {
	server := mockServer(b)
	defer server.Close()
	v2.URL = server.URL + "/"
	client := v2.NewClient()

	response, err := client.Guild(context.Background(), "Red Rose")
	assert.NoError(b, err)

	Fib := func(guildResponse *v2.GuildResponse) {
		name = guildResponse.Guild.Data.Name
	}

	for n := 0; n < b.N; n++ {
		Fib(response)
	}
}
