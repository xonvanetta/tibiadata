package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	count := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		atomic.AddInt32(&count, 1)
		cancel()
		time.Sleep(time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	err := New().Get(ctx, server.URL, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, context.Canceled))
	assert.Equal(t, int32(1), count)
}

func TestRetry(t *testing.T) {
	count := int32(0)
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if atomic.AddInt32(&count, 1) != 3 {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		writer.WriteHeader(http.StatusOK)
		_, err := writer.Write([]byte(`{"api_version":2}`))
		assert.NoError(t, err)

	}))
	defer server.Close()

	information := struct {
		APIVersion int `json:"api_version"`
	}{}
	err := New().Get(context.Background(), server.URL, &information)

	assert.NoError(t, err)
	assert.Equal(t, 2, information.APIVersion)
	assert.Equal(t, int32(3), count)
}

func TestErrorBodyText(t *testing.T) {
	response := `{"api_version":"3"}`
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte(response))
		assert.NoError(t, err)
	}))

	defer server.Close()

	information := struct {
		APIVersion int `json:"api_version"`
	}{}
	err := New().Get(context.Background(), server.URL, &information)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), response)

	errorsCount := strings.Count(err.Error(), response)
	assert.Equal(t, 1, errorsCount)
}
