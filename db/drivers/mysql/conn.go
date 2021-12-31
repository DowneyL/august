package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlConnect struct {
	username string
	password string
	host     string
	port     string
	charset  string
	dbname   string
}

func (c mysqlConnect) dns() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.username, c.password, c.host, c.port, c.dbname, c.charset,
	)
}

func NewDialector(dbname string) gorm.Dialector {
	return mysql.Open(newConnect(dbname).dns())
}

func newConnect(dbname string) mysqlConnect {
	return mysqlConnect{
		username: viper.GetString("db.mysql.username"),
		password: viper.GetString("db.mysql.password"),
		host:     viper.GetString("db.mysql.host"),
		port:     viper.GetString("db.mysql.port"),
		charset:  viper.GetString("db.mysql.charset"),
		dbname:   dbname,
	}
}
