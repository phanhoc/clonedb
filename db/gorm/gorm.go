package gorm

import (
	"github.com/phanhoc/clonedb/db"
	"github.com/jinzhu/gorm"
)

type gormDB struct {
	db     *gorm.DB
	driver string
}

func NewDB(dialect string, args ...interface{}) (db.DB, error) {
	gdb, err := gorm.Open(dialect, args...)
	if err != nil {
		return nil, err
	}
	return &gormDB{db: gdb, driver: dialect}, nil
}

func (g *gormDB) Close() {

}

func (g *gormDB) Info() {

}
