package repository

import "gorm.io/gorm"

type Init struct {
	Conexao Conexao
}

func (i *Init) Initapp() (*gorm.DB, error) {

	db, err := i.Conexao.Conn()
	if err != nil {
		return nil, err
	}

	return db, nil
}
