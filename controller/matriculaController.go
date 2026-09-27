package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/radaptech/desafio-de-vaga-jr/model"
)

type MatriculaService interface {
	Criar(matricula *m.Matricula) error
	Listar() ([]m.Matricula, error)
	BuscarPorId(id uint) (m.Matricula, error)
	BuscarComAluno(id uint) (m.Matricula, error)
	ListarPorAluno(alunoId uint) ([]m.Matricula, error)
	Atualizar(matricula *m.Matricula) error
	Deletar(id uint) error
}

type MatriculaController struct {
	service MatriculaService
}

func NewMatriculaController(service MatriculaService) *MatriculaController {
	return &MatriculaController{service: service}
}

func (ctl *MatriculaController) Criar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var matricula m.Matricula
		if err := c.ShouldBindJSON(&matricula); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		if err := ctl.service.Criar(&matricula); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusCreated, matricula)
	}
}

func (ctl *MatriculaController) Listar() gin.HandlerFunc {
	return func(c *gin.Context) {
		matriculas, err := ctl.service.Listar()
		if err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, matriculas)
	}
}

func (ctl *MatriculaController) BuscarPorId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		matricula, err := ctl.service.BuscarComAluno(id)
		if err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, matricula)
	}
}

// Usado em /alunos/:id/matriculas — o :id é o do aluno.
func (ctl *MatriculaController) ListarPorAluno() gin.HandlerFunc {
	return func(c *gin.Context) {
		alunoId, ok := parseId(c)
		if !ok {
			return
		}

		matriculas, err := ctl.service.ListarPorAluno(alunoId)
		if err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, matriculas)
	}
}

func (ctl *MatriculaController) Atualizar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		var matricula m.Matricula
		if err := c.ShouldBindJSON(&matricula); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}
		matricula.ID = id

		if err := ctl.service.Atualizar(&matricula); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, matricula)
	}
}

// PATCH: carrega a matrícula atual e sobrescreve só os campos enviados no JSON.
func (ctl *MatriculaController) AtualizarParcial() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		matricula, err := ctl.service.BuscarPorId(id)
		if err != nil {
			responderErro(c, err)
			return
		}

		if err := c.ShouldBindJSON(&matricula); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}
		matricula.ID = id

		if err := ctl.service.Atualizar(&matricula); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, matricula)
	}
}

func (ctl *MatriculaController) Deletar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		if err := ctl.service.Deletar(id); err != nil {
			responderErro(c, err)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
