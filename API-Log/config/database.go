package config

import (
	"api-log/helpers"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LocalDbConnect() (*gorm.DB, error) {
	// Open a database connection
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "your password", "19XXXXXXX", 3306, "dbname?parseTime=true")
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, helpers.ErrReturn(err)
	}
	return db, nil
}
 
