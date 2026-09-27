package service

import m "github.com/radaptech/desafio-de-vaga-jr/model"

type MatriculaRepository interface {
	AddMatricula(model *m.Matricula) error
	BuscaMatricula() ([]m.Matricula, error)
	BuscaMatriculaPorId(id uint) (m.Matricula, error)
	BuscaMatriculaComAluno(id uint) (m.Matricula, error)
	BuscaMatriculasDoAluno(alunoId uint) ([]m.Matricula, error)
	AttMatricula(model *m.Matricula) error
	DeletarMatricula(id uint) error
}

type MatriculaService struct {
	repo MatriculaRepository
}

func NewMatriculaService(repo MatriculaRepository) *MatriculaService {
	return &MatriculaService{repo: repo}
}

func (s *MatriculaService) Criar(matricula *m.Matricula) error {
	return s.repo.AddMatricula(matricula)
}

func (s *MatriculaService) Listar() ([]m.Matricula, error) {
	return s.repo.BuscaMatricula()
}

func (s *MatriculaService) BuscarPorId(id uint) (m.Matricula, error) {
	return s.repo.BuscaMatriculaPorId(id)
}

func (s *MatriculaService) BuscarComAluno(id uint) (m.Matricula, error) {
	return s.repo.BuscaMatriculaComAluno(id)
}

func (s *MatriculaService) ListarPorAluno(alunoId uint) ([]m.Matricula, error) {
	return s.repo.BuscaMatriculasDoAluno(alunoId)
}

func (s *MatriculaService) Atualizar(matricula *m.Matricula) error {
	return s.repo.AttMatricula(matricula)
}

func (s *MatriculaService) Deletar(id uint) error {
	return s.repo.DeletarMatricula(id)
}
