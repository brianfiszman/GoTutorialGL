package interfaces

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type DatabaseInterface interface {
	Connect()
	Insert(context.Context, map[string]interface{}, string) error
	Select(context.Context, []string, string) (pgx.Rows, error)
}
