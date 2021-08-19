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

type MYSQL struct {
	db *gorm.DB
}

func ConnectMYSQL() (*MYSQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &MYSQL{db}, nil
}

func CreateDatabase() {

}
