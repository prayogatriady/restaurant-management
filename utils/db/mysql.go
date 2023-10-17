package db

import (
	"fmt"
	"log"

	"github.com/prayogatriady/restaurant-management/utils/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitMySql() {

	cfg := config.AppCfg

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to MySQL: %v\n", err)
	}

	Db = db
}
