package handler

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type DownloadUrl struct {
	gorm.Model
	ID           int    `gorm:"primarykey"`
	Total_Number int    `gorm:"total_number"`
	Left_Number  int    `gorm:"left_number"`
	Used_Number  int    `gorm:"used_number"`
	DownloadUrls string `gorm:"downloadurls"`
	Used         bool   `gorm:"isused"`
	CurrentIndex int    `gorm:"currentindex"`
}

func initDB() *gorm.DB {
	dsn := "root:cw19881221@tcp(127.0.0.1:3306)/insurance"
	Db, err := gorm.Open("mysql", dsn)
	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("database link err:" + err.Error())
	}
	Db.AutoMigrate(&DownloadUrl{})
	return Db
}

func GetusedLinkNumber() string {

	database := initDB()
	// 查
	var downloadUrl DownloadUrl
	database.First(&downloadUrl, 1) // 找到id为1的产品
	return strconv.Itoa(downloadUrl.Used_Number)
}

func GetleftLinkNumber() string {
	database := initDB()
	// 查
	var downloadUrl DownloadUrl
	database.First(&downloadUrl, 1) // 找到id为1的产品
	return strconv.Itoa(downloadUrl.Left_Number)
	// return strconv.Itoa(left_number)
}

/*
func GetLastDownloadUrl() string {
	var downloadurl string
	var currrentIndex int

	database := initDB.Db

	err := database.QueryRow("select currentindex from link_numb").Scan(&currrentIndex)
	if err != nil {
		log.Panicln(err)
	}

	var isused int = 1

	for {
		err = database.QueryRow("select isused from link_numb where id = ?", currrentIndex).Scan(&isused)
		if err != nil {
			log.Panicln(err)
		}
		if isused != 1 {
			err = database.QueryRow("select download_urls from link_numb where id = ?", currrentIndex).Scan(&downloadurl)
			break
		}
		currrentIndex++
	}

	return downloadurl
}

func UpdateIndex(index int) {
	result, err := initDB.Db.Exec("UPDATE `insurance`.`link_numb` SET `isused` = '1', `currentindex` = '1' WHERE (`id` = '2')")
	if err != nil {
		log.Panicln(err)
	}
	result.LastInsertId()
}
*/
