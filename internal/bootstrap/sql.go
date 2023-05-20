package bootstrap

import (
	"fmt"
	"github.com/yonisaka/idempotency/config"
	"github.com/yonisaka/idempotency/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func RegistrySQL(conf config.DBConfig) (*gorm.DB, error) {
	var (
		dbURL string
	)

	switch conf.Driver {
	case "mysql":
		dbURL = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
		)

		db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return nil, fmt.Errorf("driver not found")
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
	)
}

func AutoSeed(db *gorm.DB) error {
	// force delete all data
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}
	// reset auto increment
	if err := db.Exec("ALTER TABLE users AUTO_INCREMENT = 1").Error; err != nil {
		return err
	}

	users := []entity.User{
		{
			Name:  "John Doe",
			Email: "johndoe@mail.com",
		},
		{
			Name:  "Yoni Saka",
			Email: "yonisaka@mail.com",
		},
	}

	if err := db.Create(&users).Error; err != nil {
		return err
	}

	return nil
}
