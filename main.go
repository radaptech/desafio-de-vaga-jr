package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/radaptech/desafio-de-vaga-jr/controller"
	"github.com/radaptech/desafio-de-vaga-jr/repository"
	"github.com/radaptech/desafio-de-vaga-jr/service"
)

func main() {

	init := repository.Init{Conexao: &repository.ConnTestSqlite{}}
	db, err := init.Initapp()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	alunoCtl := controller.NewAlunoController(service.NewAlunoService(repo))
	matriculaCtl := controller.NewMatriculaController(service.NewMatriculaService(repo))

	r := gin.Default()

	alunos := r.Group("/alunos")
	alunos.POST("", alunoCtl.Criar())
	alunos.GET("", alunoCtl.Listar())
	alunos.GET("/:id", alunoCtl.BuscarPorId())
	alunos.GET("/:id/matriculas", matriculaCtl.ListarPorAluno())
	alunos.PUT("/:id", alunoCtl.Atualizar())
	alunos.PATCH("/:id", alunoCtl.AtualizarParcial())
	alunos.DELETE("/:id", alunoCtl.Deletar())

	matriculas := r.Group("/matriculas")
	matriculas.POST("", matriculaCtl.Criar())
	matriculas.GET("", matriculaCtl.Listar())
	matriculas.GET("/:id", matriculaCtl.BuscarPorId())
	matriculas.PUT("/:id", matriculaCtl.Atualizar())
	matriculas.PATCH("/:id", matriculaCtl.AtualizarParcial())
	matriculas.DELETE("/:id", matriculaCtl.Deletar())

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
