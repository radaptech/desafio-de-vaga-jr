package repository

import (
	"errors"
	"fmt"

	m "github.com/radaptech/desafio-de-vaga-jr/model"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {

	return &Repository{
		Db: db,
	}
}

func (r *Repository) AddAluno(model *m.Aluno) error {

	result := r.Db.Create(model)

	return result.Error
}

func (r *Repository) BuscaAluno() ([]m.Aluno, error) {

	var alunos []m.Aluno
	result := r.Db.Find(&alunos)
	if result.Error != nil {
		return nil, result.Error
	}

	return alunos, nil
}

func (r *Repository) BuscaAlunoPorId(id uint) (m.Aluno, error) {

	var Aluno m.Aluno

	result := r.Db.First(&Aluno, id)
	if result.Error != nil {
		return m.Aluno{}, result.Error
	}

	return Aluno, nil
}

func (r *Repository) BuscaAlunoMatricula() ([]m.Aluno, error) {

	var alunos []m.Aluno

	result := r.Db.Preload("Matriculas").Find(&alunos)
	if result.Error != nil {
		return nil, result.Error
	}

	return alunos, nil
}

func (r *Repository) BuscaAlunoPorIdMatricula(id uint) (m.Aluno, error) {

	var Aluno m.Aluno

	result := r.Db.Preload("Matriculas").First(&Aluno, id)
	if result.Error != nil {
		return m.Aluno{}, result.Error
	}

	return Aluno, nil
}

func (r *Repository) AttAlunos(model *m.Aluno) error {

	var existe m.Aluno

	if err:= r.Db.First(&existe, model.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("aluno não encontrado: %w", gorm.ErrRecordNotFound)
		}
		return err
	}

	result:= r.Db.Save(model)
	return  result.Error
}

func (r *Repository) AttAlunoCampo(id uint, dados map[string]any) error {

	result:= r.Db.Model(&m.Aluno{}).Where("id = ?", id).Updates(dados)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("aluno nao encontrado: %w", gorm.ErrRecordNotFound)
	}

	return  nil
}

func (r *Repository) DelatarAluno(id uint) error {

	var aluno m.Aluno

	if err := r.Db.First(&aluno, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return fmt.Errorf("aluno não encontrado: %w", gorm.ErrRecordNotFound)
		}

		return err
	}

	result:= r.Db.Delete(&aluno)
	return result.Error
}