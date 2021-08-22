package dbsource

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USER     = "root"
	PASSWORD = "123456"
	IP       = "127.0.0.1"
	DBNAME   = "ituring"
)

func ConnectMYSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASSWORD, IP, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}
	// Ping
	err = sqlDB.Ping()

	if err != nil {
		sqlDB.Close() // Close
		return nil, err
	}

	return db, nil
}
