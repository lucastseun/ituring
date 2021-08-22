package dbsource

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"
	password = "123456"
	ip       = "127.0.0.1"
	dbname   = "ituring"
)

func ConnectMYSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, dbname)
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
