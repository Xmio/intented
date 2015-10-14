package main

import (
	"log"
	"net/http"

	"database/sql"

	"github.com/GeertJohan/go.rice"
	"github.com/Xmio/intented/server/datastores"
	"github.com/Xmio/intented/server/lead"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

type config struct {
	Port        string `env:"PORT" envDefault:"3000"`
	Production  bool   `env:"PRODUCTION"`
	DatabaseURL string `env:"DATABASE_URL" envDefault:"postgres://localhost:5432/intented?sslmode=disable"`
}

func main() {
	var config config
	env.Parse(&config)
	log.Println(config)
	db := datastores.NewDBConnectionPool(config.DatabaseURL)
	defer db.Close()
	exec := server(config, db)
	exec.Run(":" + config.Port)
}

func server(config config, db *sql.DB) *echo.Echo {
	dbx := sqlx.NewDb(db, "postgres")
	exec := echo.New()
	if !config.Production {
		exec.Debug()
	}
	exec.Use(mw.Logger())
	exec.Use(mw.Recover())
	exec.Use(mw.Gzip())

	exec.Get("/status", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	leadHandler := lead.NewHandler(datastores.NewLead(dbx))

	exec.Post("/lead", leadHandler.Create)
	exec.Get("/lead/:hashCode", leadHandler.CountByInvites)

	assetHandler := http.FileServer(rice.MustFindBox("static").HTTPBox())
	exec.Get("/", func(c *echo.Context) error {
		assetHandler.ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})
	exec.Get("/static/*", func(c *echo.Context) error {
		http.StripPrefix("/static/", assetHandler).
			ServeHTTP(c.Response().Writer(), c.Request())
		return nil
	})

	return exec
}
