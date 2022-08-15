package initdb

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	dsn := "root:cw19881221@tcp(127.0.0.1:3306)/insurance"
	Db, err := gorm.Open("mysql", dsn)
	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("database link err:" + err.Error())
	}
	defer Db.Close()
}
