package main

import (
	"bufio"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Xmio/intented/datastores"
	"github.com/caarlos0/env"
	"github.com/jmoiron/sqlx"
)

type testConfig struct {
	CreateDB     bool   `env:"CREATE_TEST_DB" envDefault:"true"`
	MigrateDB    bool   `env:"MIGRATE_DB" envDefault:"true"`
	PostgresURL  string `env:"POSTGRES_URL" envDefault:"postgres://localhost:5432"`
	DatabaseName string
}

func destroyTestDB(db *sqlx.DB, testCfg testConfig) {
	db.Close()
	if testCfg.CreateDB {
		pgExec("DROP DATABASE "+testCfg.DatabaseName, testCfg)
	}
}

func pgExec(stm string, testCfg testConfig) {
	db, err := sqlx.Connect("postgres", testCfg.PostgresURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if _, err = db.Exec(stm); err != nil {
		panic(err)
	}
}

func createTestDatabase(testCfg testConfig) string {
	name := randomStr()
	log.Println("Create-ing test database " + name)
	pgExec("CREATE DATABASE "+name, testCfg)
	return name
}

func randomStr() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 10)
	for i := 0; i < 10; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func prepareTestConfig() (config, testConfig) {
	var cfg config
	var testCfg testConfig
	env.Parse(&cfg)
	env.Parse(&testCfg)
	if testCfg.CreateDB {
		testCfg.DatabaseName = createTestDatabase(testCfg)
		cfg.DatabaseURL = buildDBURL(testCfg)
	}
	log.Println(cfg)
	log.Println(testCfg)
	return cfg, testCfg
}

func buildDBURL(cfg testConfig) string {
	pgURL := cfg.PostgresURL
	if strings.HasSuffix(pgURL, "/") {
		pgURL = pgURL[:len(pgURL)-1]
	}
	if strings.Contains(pgURL, "?") {
		pgURL := strings.Split(pgURL, "?")
		return pgURL[0] + "/" + cfg.DatabaseName + "?" + pgURL[1]
	}
	return pgURL + "/" + cfg.DatabaseName
}

func prepareTestDB(cfg config, testCfg testConfig) *sqlx.DB {
	db := datastores.NewDBConnectionPool(cfg.DatabaseURL)
	if testCfg.MigrateDB {
		migrate(db)
	}
	return db
}

func migrate(db *sqlx.DB) {
	log.Println("Migrate-ing database...")
	files, _ := filepath.Glob(filepath.Join("../migrations", "*.sql"))
	for _, file := range files {
		log.Println("Migrating file", file)
		file, _ := os.Open(file)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		sql := ""
		for scanner.Scan() {
			sql += scanner.Text()
		}
		if _, err := db.Exec(sql); err != nil {
			panic(err)
		}
	}
}

func putForm(url string, data url.Values) (resp *http.Response, err error) {
	return put(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func put(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	client := http.Client{}
	return client.Do(req)
}
