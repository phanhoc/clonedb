package db

import "github.com/phanhoc/clonedb/model/sun"

type DB interface {
	Close() error
	Info()
	MigrateSchema() error
	InsertNiche(*sun.TShirt) error
}
