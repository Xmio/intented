package intented

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/labstack/echo"
)

var ErrExpected = errors.New("Expected Error")

func request(method, path string, data url.Values, e *echo.Echo) (int, string) {
	if data == nil {
		data = url.Values{}
	}
	r, _ := http.NewRequest(method, path, bytes.NewBufferString(data.Encode()))
	r.Header.Add(echo.ContentType, echo.ApplicationForm)
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	return w.Code, w.Body.String()
}
