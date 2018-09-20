package Models

import (
	"log"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = os.Getenv("dbType")
	dbName = os.Getenv("dbName")
	user = os.Getenv("user")
	password = os.Getenv("password")
	host = os.Getenv("host")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
}

func CloseDB() {
	defer db.Close()
}
