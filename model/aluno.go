package aluno

import (
	"time"

	"gorm.io/gorm"
)

type Matricula struct {
	gorm.Model
	CodigoMatricula string    `gorm:"type:varchar(20);not null" json:"codigoMatricula"`
	NomeCurso       string    `gorm:"type:varchar(20);not  null" json:"nomeCurso"`
	DataInicio      time.Time `gorm:"type:date;not null" json:"dataInicio"`
	AlunoId         int       `gorm:"not null" json:"alunoId"`
	Aluno           Aluno     `gorm:"foreignKey:AlunoId; constraint:OnDelete:CASCADE"`
}

type Aluno struct {
	gorm.Model
	Nome           string    `gorm:"type:varchar(30); not null" json:"nome"`
	Telefone       string    `gorm:"type:varchar(15);not null" json:"telefone"`
	DataNascimento time.Time `gorm:"type:date;not null" json:"dataNascimento"`
	Matriculas     []Matricula
}
