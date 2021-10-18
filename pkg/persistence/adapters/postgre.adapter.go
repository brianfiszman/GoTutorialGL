package adapters

import (
	"context"
	"fmt"
	"os"
	"training/tutorialGL/pkg/persistence/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	Config         config.DatabaseConfig
	ConnectionPool *pgxpool.Pool
}

func (d *PostgreSQL) Connect() {
	dbpool, err := pgxpool.Connect(context.Background(), d.Config.GetConnectionString())

	d.ConnectionPool = dbpool

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to PostgreSQL")
}
