package repository

import (
	"errors"
	"fmt"

	m "github.com/radaptech/desafio-de-vaga-jr/model"
	"gorm.io/gorm"
)

func (r *Repository) AddMatricula(model *m.Matricula) error {

	result := r.Db.Create(model)

	return result.Error
}

func (r *Repository) BuscaMatricula() ([]m.Matricula, error) {

	var matriculas []m.Matricula
	result := r.Db.Find(&matriculas)
	if result.Error != nil {
		return nil, result.Error
	}

	return matriculas, nil
}

func (r *Repository) BuscaMatriculaPorId(id uint) (m.Matricula, error) {

	var matricula m.Matricula

	result := r.Db.First(&matricula, id)
	if result.Error != nil {
		return m.Matricula{}, result.Error
	}

	return matricula, nil
}

func (r *Repository) BuscaMatriculaAluno() ([]m.Matricula, error) {

	var matriculas []m.Matricula

	result := r.Db.Preload("Aluno").Find(&matriculas)
	if result.Error != nil {
		return nil, result.Error
	}

	return matriculas, nil
}

func (r *Repository) BuscaMatriculaComAluno(id uint) (m.Matricula, error) {

	var matricula m.Matricula

	result := r.Db.Preload("Aluno").First(&matricula, id)
	if result.Error != nil {
		return m.Matricula{}, result.Error
	}

	return matricula, nil
}

func (r *Repository) BuscaMatriculasDoAluno(alunoId uint) ([]m.Matricula, error) {

	var matriculas []m.Matricula

	result := r.Db.Preload("Aluno").Where("aluno_id = ?", alunoId).Find(&matriculas)
	if result.Error != nil {
		return nil, result.Error
	}

	return matriculas, nil
}

func (r *Repository) AttMatricula(model *m.Matricula) error {

	var existe m.Matricula

	if err := r.Db.First(&existe, model.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("matricula não encontrada: %w", gorm.ErrRecordNotFound)
		}
		return err
	}

	result := r.Db.Save(model)
	return result.Error
}

func (r *Repository) AttMatriculaCampo(id uint, dados map[string]any) error {

	result := r.Db.Model(&m.Matricula{}).Where("id = ?", id).Updates(dados)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("matricula nao encontrada: %w", gorm.ErrRecordNotFound)
	}

	return nil
}

func (r *Repository) DeletarMatricula(id uint) error {

	var matricula m.Matricula

	if err := r.Db.First(&matricula, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("matricula não encontrada: %w", gorm.ErrRecordNotFound)
		}

		return err
	}

	result := r.Db.Delete(&matricula)
	return result.Error
}
