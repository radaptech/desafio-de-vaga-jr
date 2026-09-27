package service

import m "github.com/radaptech/desafio-de-vaga-jr/model"

type AlunoRepository interface {
	AddAluno(model *m.Aluno) error
	BuscaAluno() ([]m.Aluno, error)
	BuscaAlunoPorId(id uint) (m.Aluno, error)
	BuscaAlunoPorIdMatricula(id uint) (m.Aluno, error)
	AttAlunos(model *m.Aluno) error
	DelatarAluno(id uint) error
}

type AlunoService struct {
	repo AlunoRepository
}

func NewAlunoService(repo AlunoRepository) *AlunoService {
	return &AlunoService{repo: repo}
}

func (s *AlunoService) Criar(aluno *m.Aluno) error {
	return s.repo.AddAluno(aluno)
}

func (s *AlunoService) Listar() ([]m.Aluno, error) {
	return s.repo.BuscaAluno()
}

func (s *AlunoService) BuscarPorId(id uint) (m.Aluno, error) {
	return s.repo.BuscaAlunoPorId(id)
}

func (s *AlunoService) BuscarComMatriculas(id uint) (m.Aluno, error) {
	return s.repo.BuscaAlunoPorIdMatricula(id)
}

func (s *AlunoService) Atualizar(aluno *m.Aluno) error {
	return s.repo.AttAlunos(aluno)
}

func (s *AlunoService) Deletar(id uint) error {
	return s.repo.DelatarAluno(id)
}
