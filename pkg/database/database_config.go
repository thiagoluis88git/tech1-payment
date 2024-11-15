package database

import (
	"github.com/thiagoluis88git/tech1-payment/internal/core/data/model"

	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

func ConfigDatabase(dialector gorm.Dialector) (*Database, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		return &Database{}, err
	}

	db.AutoMigrate(
		&model.Order{},
		&model.OrderProduct{},
		&model.Payment{},
	)

	return &Database{
		Connection: db,
	}, nil
}
