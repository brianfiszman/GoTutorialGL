package adapters

import (
	"context"
	"fmt"
	"os"
	"training/tutorialGL/pkg/persistence/config"

	sq "github.com/Masterminds/squirrel"
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

func (d *PostgreSQL) Insert(ctx context.Context, entity map[string]interface{}, entityName string) error {
	keys := make([]string, 0, len(entity))
	values := make([]interface{}, 0, len(entity))

	for k, v := range entity {
		keys = append(keys, k)
		values = append(values, v)
	}

	fmt.Println("KEYS", keys)
	fmt.Println("VALUES", values)

	query, _, _ := sq.Insert(entityName).Columns(keys...).Values(values...).PlaceholderFormat(sq.Dollar).ToSql()

	fmt.Println("QUERYSTRING", query)

	if _, err := d.ConnectionPool.Exec(context.Background(), query, values...); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return err
	}

	return nil
}
