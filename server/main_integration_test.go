package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func TestMain(m *testing.M) {
	cfg, testCfg := prepareTestConfig()
	db := prepareTestDB(cfg, testCfg)
	defer destroyTestDB(db, testCfg)
	echoTest := server(cfg, db)
	testServer = httptest.NewServer(http.HandlerFunc(echoTest.ServeHTTP))
	defer testServer.Close()
	m.Run()
}

func TestHealthCheck(t *testing.T) {
	resp, err := http.Get(testServer.URL + "/status")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
