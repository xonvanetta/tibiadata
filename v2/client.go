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
	URL     = "https://api.tibiadata.com/v2/"
	Retries = 5
)

var (
	location               *time.Location
	jsonUnmarshalTypeError *json.UnmarshalTypeError
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

type Errors []*Error

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

func (e Errors) IsAllNotFound() bool {
	allNotFound := true
	for _, err := range e {
		if !err.IsNotFound() {
			allNotFound = false
			break
		}
	}

	return allNotFound
}

func (e Errors) Unwrap() error {
	return e[len(e)-1]
}

func (e *Errors) Add(err *Error) {
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

type Error struct {
	Err        error
	StatusCode int
	Url        string
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Error() string {
	return fmt.Sprintf("url: %s, statusCode: %d, err: %s", e.Url, e.StatusCode, e.Err)
}

func (e Error) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

func requestError(request *http.Request, err error) *Error {
	e := &Error{
		Err: err,
		Url: request.URL.String(),
	}
	if request.Response != nil {
		e.StatusCode = request.Response.StatusCode
	}
	return e
}

func (c Client) request(request *http.Request, v interface{}) *Error {
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
