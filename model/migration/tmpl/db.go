package tmpl

var DBTemplate = `package db

import (
	"github.com/DowneyL/august/db"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

var (
	dbConnect *gorm.DB
	once      sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		dbConnect = db.Use(viper.GetString("{{.AppName}}.db_name"))
	})

	return dbConnect
}
`
