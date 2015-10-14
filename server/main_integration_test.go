package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Xmio/intented/server/datastores"
	"github.com/caarlos0/it"
	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	serverFn := func(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
		return server(config{}, db).ServeHTTP
	}
	it := it.New()
	handler := it.Init(serverFn, datastores.NewDBConnectionPool)
	defer it.Shutdown()
	testServer = httptest.NewServer(http.HandlerFunc(handler))
	defer testServer.Close()
	m.Run()
}

func TestStatus(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/status")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
