package aluno

import (
	"time"

	"gorm.io/gorm"
)

type Matriculas struct {

	gorm.Model
	Id int `gorm:"primaryKey" json:"id"`
	CodigoMatricula string `gorm:"varchar(20);not null" json:"codigoMatricula"`
	NomeCurso string `gorm:"varchar(20);not  null" json:"nomeCurso"`
	DataInicio time.Time `gorm:"type:date;not null" json:"dataInicio"`
	Aluno Aluno
}


type Aluno struct {
	Id              int
	Nome            string
	Telefone        string
	DataNascimento  time.Time
	DataInclusao    time.Time
	DataAtualizacao time.Time
	Matriculas      []Matriculas
}
