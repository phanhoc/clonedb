package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phanhoc/clonedb/db"
	"github.com/phanhoc/clonedb/model/sun"
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

func (g *gormDB) Close() error {
	return g.db.Close()
}

func (g *gormDB) Info() {

}

func (g *gormDB) MigrateSchema() error {
	if err := g.db.AutoMigrate(
		&sun.TShirt{}).Error; err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}

func (g *gormDB) InsertNiche(niche *sun.TShirt) error {
	return g.db.Create(niche).Error
}
