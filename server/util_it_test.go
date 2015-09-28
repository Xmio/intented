package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatesDBURLWithoutSSLMode(t *testing.T) {
	cfg := testConfig{
		PostgresURL:  "postgres://localhost:5432",
		DatabaseName: "test",
	}
	url := buildDBURL(cfg)
	assert.Equal(t, url, "postgres://localhost:5432/test")
}

func TestCreatesDBURLWithSSLMode(t *testing.T) {
	cfg := testConfig{
		PostgresURL:  "postgres://localhost:5432?sslmode=disable",
		DatabaseName: "test",
	}
	url := buildDBURL(cfg)
	assert.Equal(t, url, "postgres://localhost:5432/test?sslmode=disable")
}

func TestCreatesDBURLWithTrailingSlash(t *testing.T) {
	cfg := testConfig{
		PostgresURL:  "postgres://localhost:5432/",
		DatabaseName: "test",
	}
	url := buildDBURL(cfg)
	assert.Equal(t, url, "postgres://localhost:5432/test")
}
