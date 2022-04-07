package database

import (
	"log"
	"rest-api-with-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBanco() {
	dsn := "host=localhost user=root password=root dbname=alunos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Erro ao conectar com banco de dados.")
	}

	DB.AutoMigrate(&models.Aluno{})
}
