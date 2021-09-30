package httpclient

import (
	"fmt"
	"net/http"
)

type Errors []*Error

var (
	ErrWrongStatusCode = fmt.Errorf("wrong status code")
)

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

func (e *Errors) Add(err *Error) {
	*e = append(*e, err)
}

func responseError(response *http.Response, err error) *Error {
	return &Error{
		Err:        err,
		Url:        response.Request.URL.String(),
		StatusCode: response.StatusCode,
	}
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
