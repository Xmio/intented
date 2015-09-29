package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
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

func TestCreateLead(t *testing.T) {
	mail := "jem@intented.co"
	resp, err := http.PostForm(testServer.URL+"/leads/"+mail, url.Values{})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateLeadWithInvited(t *testing.T) {
	mail := "jem@intented.co"
	invitedCode := "2d2258672454bc08c6c417f8b390da7a"
	resp, err := http.PostForm(testServer.URL+"/leads/"+mail+"/"+invitedCode, url.Values{})
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func CountLeadByInvit(t *testing.T) {
	mail := "jem@intented.co"
	mail2 := "jem2@intented.co"
	invitedCode := "2d2258672454bc08c6c417f8b390da7a"
	countURL := testServer.URL + "/lead/" + invitedCode
	http.PostForm(testServer.URL+"/leads/"+mail+"/"+invitedCode, url.Values{})
	resp, err := http.Get(countURL)
	assert.NoError(t, err)
	count := toInt(resp.Body)
	assert.Equal(t, count, 1)
	http.PostForm(testServer.URL+"/leads/"+mail2+"/"+invitedCode, url.Values{})
	resp2, err := http.Get(countURL)
	assert.NoError(t, err)
	count2 := toInt(resp2.Body)
	assert.Equal(t, count2, 2)
}

func toInt(body io.ReadCloser) int64 {
	var result int64
	bts, _ := ioutil.ReadAll(body)
	json.Unmarshal(bts, &result)
	return result
}
