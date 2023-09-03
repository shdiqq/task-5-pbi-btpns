package config

import (
	"fmt"
	"log"

	"github.com/shdiqq/task-5-pbi-btpns-Shadiq/models/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbConnection := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%vJakarta", ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_HOST, ENV.DB_PORT, ENV.DB_DATABASE, "%2F")
	database, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	if err != nil {
		log.Println("[ERROR] Gagal koneksi database", ENV.DB_DATABASE)
		log.Panic(err)
	}
	database.AutoMigrate(&entity.User{}, &entity.Photo{})

	DB = database
	log.Println("[Database] Berhasil koneksi database", ENV.DB_DATABASE)
}
