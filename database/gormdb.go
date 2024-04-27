package database

import (
	"crud-customer/config"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GormDB interface {
	GetDB() *gorm.DB
	AutoMigrate() error
}

type gormDB struct {
	db  *gorm.DB
	cfg *config.Config
}

func (g *gormDB) GetDB() *gorm.DB {
	return g.db
}

func NewGormDB(cfg *config.Config) (GormDB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.File), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &gormDB{db: db, cfg: cfg}, nil
}
