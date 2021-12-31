package db

import (
	"github.com/DowneyL/august/db/drivers/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

func Use(dbname string) *gorm.DB {
	return open(newDialector(dbname))
}

func open(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, getOpts())

	if err != nil {
		panic(err)
	}

	if sqlDb, err2 := db.DB(); err2 == nil {
		sqlDb.SetMaxIdleConns(10)
		sqlDb.SetMaxOpenConns(100)
		sqlDb.SetConnMaxLifetime(time.Hour)
	}

	return db
}

func getOpts() *gorm.Config {
	return &gorm.Config{}
}

func newDialector(dbname string) gorm.Dialector {
	switch viper.GetString("db_connection") {
	case "mysql":
		return mysql.NewDialector(dbname)
	default:
		panic("unregister db connection")
	}
}
