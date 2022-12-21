package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	connectString := "root:@/simpletest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
