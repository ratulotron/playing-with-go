package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func LoadDB() {

}

var db *gorm.DB

func init() {

	connString := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
	)

	newDB, err := gorm.Open("mysql", connString)
	if err != nil {
		os.Exit(1)
	}

	db = newDB
}

func GetConnection() *gorm.DB {
	return db
}
