package clients

import (
	"backend/dao"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() error {
	dsn := "root:root@tcp(localhost:3306)/coursesPlatform?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	err = db.AutoMigrate(&dao.User{}, &dao.Course{}, &dao.Subscription{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}

func SelectUserByID(id int64) (dao.User, error) {
	var user dao.User
	result := db.First(&user, id)
	if result.Error != nil {
		return dao.User{}, fmt.Errorf("not found user with ID: %d", id)
	}
	return user, nil
}
