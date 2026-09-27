package repository

import (
	"github.com/radaptech/desafio-de-vaga-jr/model"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)


type Conexao interface {
	Conn()(*gorm.DB, error)
}

type ConnTestSqlite struct {}
func (c *ConnTestSqlite )Conn() (*gorm.DB, error) {
	
	db, err:= gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return  nil, err
	}

	sqlDb, err:= db.DB()
	if err != nil {
		return  nil, err
	} 

	sqlDb.SetMaxOpenConns(1)

	if err := db.AutoMigrate(&aluno.Aluno{}, &aluno.Matricula{}); err != nil {

		return nil, err
	}
	return db, nil
}


