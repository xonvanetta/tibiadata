package integration

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	v2 "github.com/xonvanetta/tibiadata/v2"
)

func info(format string, args ...interface{}) {
	log.Println(fmt.Sprintf("Integration Test: "+format, args...))
}

func TestAll(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test, Really long test is not verified yet. JSON unmarshal error guaranteed.")
	}
	client := v2.New()

	worlds, err := client.Worlds(context.Background())
	assert.NoError(t, err)

	for _, world := range worlds.Worlds.Allworlds {
		info("running world: %s", world.Name)
		_, err := client.World(context.Background(), world.Name)
		assert.NoError(t, err)

		guilds, err := client.Guilds(context.Background(), world.Name)
		assert.NoError(t, err)

		for _, guild := range guilds.Guilds.Active {
			info("running guild: %s, world: %s", guild.Name, world.Name)
			_, err := client.Guild(context.Background(), guild.Name)
			assert.NoError(t, err)
		}
	}
}

//func TestAll(t *testing.T) {
//	client := v2.New()
//
//	ctx, cancel := context.WithCancel(context.Background())
//
//	ch, done := workers(ctx, 100)
//	errCh := make(chan error)
//
//	sendError := func(err error) {
//		if err == nil {
//			return
//		}
//		errCh <- err
//	}
//
//	worlds, err := client.Worlds(context.Background())
//	assert.NoError(t, err)
//
//	for _, world := range worlds.Worlds.Allworlds {
//		world := world
//		ch <- func() {
//			ctx, cancel := context.WithCancel(context.Background())
//			ch, done := workers(ctx, 10)
//			info("running world: %s", world.Name)
//			_, err := client.World(context.Background(), world.Name)
//			sendError(err)
//
//			guilds, err := client.Guilds(context.Background(), world.Name)
//			sendError(err)
//
//			for _, guild := range guilds.Guilds.Active {
//				guild := guild
//				ch <- func() {
//					info("running guild: %s, world: %s", guild.Name, world.Name)
//					_, err := client.Guild(context.Background(), guild.Name)
//					sendError(err)
//				}
//			}
//			cancel()
//			<-done
//		}
//	}
//	cancel()
//	select {
//	case err := <-errCh:
//		t.Error(err)
//	case <-done:
//		fmt.Println("done")
//	}
//}
//
//func workers(ctx context.Context, amount int) (chan func(), chan struct{}) {
//	workerCh := make(chan func())
//	done := make(chan struct{})
//
//	wg := sync.WaitGroup{}
//	wg.Add(amount)
//
//	for i := 0; i < amount; i++ {
//		go func() {
//			select {
//			case callback := <-workerCh:
//				callback()
//			case <-ctx.Done():
//				wg.Done()
//				return
//			}
//		}()
//	}
//
//	go func() {
//		wg.Wait()
//		close(done)
//	}()
//
//	return workerCh, done
//}
