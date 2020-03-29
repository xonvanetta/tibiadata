package v2

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
	URL     = "http://api.tibiadata.com/v2/"
	Retries = 5
)

type Client struct {
	http *http.Client
}

func New() Client {
	return Client{
		http: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewWithClient(client *http.Client) Client {
	return Client{
		http: client,
	}
}

type Errors []error

func (e Errors) Error() string {
	text := ""
	for i, err := range e {
		if i == 0 {
			text = fmt.Sprintf("error %d: %s", i, err)
			continue
		}
		text = fmt.Sprintf("%s\nerror %d: %s", text, i, err)
	}
	return text
}

func (e Errors) Unwrap() error {
	return e[len(e)-1]
}

func (e *Errors) Add(err error) {
	*e = append(*e, err)
}

func (c Client) get(context context.Context, path string, v interface{}) error {
	url := fmt.Sprintf("%s%s", URL, path)

	request, err := http.NewRequestWithContext(context, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request for url: %s, err: %w", url, err)
	}

	var errs Errors

	for i := 0; i < Retries; i++ {
		select {
		case <-context.Done():
			return errs
		default:
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

var jsonUnmarshalTypeError *json.UnmarshalTypeError

func requestError(request *http.Request, err error) error {
	if request.Response == nil {
		return fmt.Errorf("url: %s, err: %w", request.URL, err)
	}
	return fmt.Errorf("url: %s, statusCode: %d, err: %w", request.URL, request.Response.StatusCode, err)
}

func (c Client) request(request *http.Request, v interface{}) error {
	response, err := c.http.Do(request)
	if err != nil {
		return requestError(request, fmt.Errorf("failed to do request: %w", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return requestError(request, fmt.Errorf("wrong status code"))
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return requestError(request, fmt.Errorf("failed to read body: %w", err))
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return requestError(request, fmt.Errorf("failed to decode request: '%s', err: %w", string(data), err))
	}
	return nil
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
	location, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {
		panic(fmt.Errorf("failed to load location for tibiadata: %w", err))
	}

	t.Time, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, string(b), location)
	return err
}
