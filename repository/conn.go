package repository

import (
	"github.com/radaptech/desafio-de-vaga-jr/model"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)


func Conn() (*gorm.DB, error) {
	
	db, err:= gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return  nil, err
	}


	if err := db.AutoMigrate(&aluno.Aluno{}, &aluno.Matriculas{}); err != nil {

		return nil, err
	}
	return db, nil
}