package interfaces

import "context"

type DatabaseInterface interface {
	Connect()
	Insert(context.Context, map[string]interface{}, string) error
}
