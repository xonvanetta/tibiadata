package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	jsonUnmarshalTypeError *json.UnmarshalTypeError
	UserAgent              = "tibiadata/v4"
)

type Client struct {
	http    *http.Client
	retries int
}

func New() Client {
	return Client{
		http: &http.Client{
			Timeout: time.Second * 30,
		},
		retries: 5,
	}
}

func NewWithClient(client *http.Client, retries int) Client {
	return Client{
		http:    client,
		retries: retries,
	}
}

func (c Client) Get(context context.Context, url string, v interface{}) error {
	request, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request for url: %s, err: %w", url, err)
	}

	request.Header.Set("user-agent", UserAgent)

	var errs Errors

	for i := 0; i < c.retries; i++ {
		if context.Err() != nil {
			errs.Add(&Error{
				Err: context.Err(),
				Url: request.URL.String(),
			})
			return errs
		}
		err := c.request(request, v)
		if err != nil {
			errs.Add(err)
			if errors.As(err, &jsonUnmarshalTypeError) {
				return errs
			}
			continue
		}
		return nil
	}
	return errs
}

func (c Client) request(request *http.Request, v interface{}) *Error {
	response, err := c.http.Do(request)
	if err != nil {
		return &Error{
			Err: err,
			Url: request.URL.String(),
		}
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return responseError(response, ErrWrongStatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseError(response, fmt.Errorf("failed to read body: %w", err))
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return responseError(response, fmt.Errorf("failed to decode request: '%s', err: %w", string(data), err))
	}
	return nil
}
