package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "github.com/radaptech/desafio-de-vaga-jr/model"
)

type AlunoService interface {
	Criar(aluno *m.Aluno) error
	Listar() ([]m.Aluno, error)
	BuscarPorId(id uint) (m.Aluno, error)
	BuscarComMatriculas(id uint) (m.Aluno, error)
	Atualizar(aluno *m.Aluno) error
	Deletar(id uint) error
}

type AlunoController struct {
	service AlunoService
}

func NewAlunoController(service AlunoService) *AlunoController {
	return &AlunoController{service: service}
}

func (ctl *AlunoController) Criar() gin.HandlerFunc {
	return func(c *gin.Context) {
		var aluno m.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		if err := ctl.service.Criar(&aluno); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusCreated, aluno)
	}
}

func (ctl *AlunoController) Listar() gin.HandlerFunc {
	return func(c *gin.Context) {
		alunos, err := ctl.service.Listar()
		if err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, alunos)
	}
}

func (ctl *AlunoController) BuscarPorId() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		aluno, err := ctl.service.BuscarComMatriculas(id)
		if err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, aluno)
	}
}

func (ctl *AlunoController) Atualizar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		var aluno m.Aluno
		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}
		aluno.ID = id

		if err := ctl.service.Atualizar(&aluno); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, aluno)
	}
}

// PATCH: carrega o aluno atual e sobrescreve só os campos enviados no JSON.
func (ctl *AlunoController) AtualizarParcial() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := parseId(c)
		if !ok {
			return
		}

		aluno, err := ctl.service.BuscarPorId(id)
		if err != nil {
			responderErro(c, err)
			return
		}

		if err := c.ShouldBindJSON(&aluno); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}
		aluno.ID = id

		if err := ctl.service.Atualizar(&aluno); err != nil {
			responderErro(c, err)
			return
		}

		c.JSON(http.StatusOK, aluno)
	}
}

func (ctl *AlunoController) Deletar() gin.HandlerFunc {
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
