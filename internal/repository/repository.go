package repository

import (
	"fmt"
	"question-answer-service/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepository(config config.ConfigDatabase) (*repository, error) {
	dsn := buildDatabaseDSN(config)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Printf("Error during GORM DB initialization: %v\n", err)
		return &repository{}, err
	}
	return &repository{
		DB: db,
	}, nil
}

func buildDatabaseDSN(configDB config.ConfigDatabase) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configDB.Host,
		configDB.User,
		configDB.Password,
		configDB.Name,
		configDB.Port,
	)
}
