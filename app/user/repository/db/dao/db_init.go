package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gorm.io/gorm/logger"
	"micro-todoList/micro-todoList/config"
	"strings"
)

var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	port := config.DbPort
	database := config.DbName
	username := config.DbUser
	password := config.DbPassWord
	charset := config.Charset
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func Database(connString string) error {
	var ormLogger logger.Interface

}
