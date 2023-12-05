package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/MarselBissengaliyev/bank/config"
	"github.com/MarselBissengaliyev/bank/db/postgres"
	"github.com/MarselBissengaliyev/bank/internal/repo"
)

func TestMain(m *testing.M) {
	postgresCfg, err := config.InitPostgresConfig("../../config/")

	if err != nil {
		log.Fatal("Failed to initialize postgres config: ", err)
	}

	conn, err := postgres.NewPostgresDB(postgresCfg)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	repo.NewRepository(conn)

	os.Exit(m.Run())
}
