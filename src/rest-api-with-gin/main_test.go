package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"rest-api-with-gin/controllers"
	"rest-api-with-gin/database"
	"rest-api-with-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTestes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno de teste",
		CPF:  "1234567890",
		RG:   "9876543210",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	aluno := models.Aluno{}
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDoEndpointHelloComParametro(t *testing.T) {
	r := SetupRotasTestes()
	r.GET("/hello/:nome", controllers.Hello)

	req, _ := http.NewRequest("GET", "/hello/Felipe", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Error: O status da requisição foi %d e o esperado era %d.", res.Code, http.StatusOK)

	mockDaResposta := `{"data":"Fala Felipe! Mais um pateta pro Go. Welcome bro!"}`

	corpoDaRequisicao, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, mockDaResposta, string(corpoDaRequisicao), "Error: O retorno da requisição é diferente do esperado.")
}

func TestListandoTodosOsAluno(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTestes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Error: O status da requisição foi %d e o esperado era %d.", res.Code, http.StatusOK)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTestes()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/1234567890", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "Error: O status da requisição foi %d e o esperado era %d.", res.Code, http.StatusOK)
}

func TestBuscaAlunoPorID(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTestes()
	r.GET("/alunos/:id", controllers.ExibeAluno)

	urlDaBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", urlDaBusca, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno de teste", alunoMock.Nome)
	assert.Equal(t, "1234567890", alunoMock.CPF)
	assert.Equal(t, "9876543210", alunoMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeletaAluno(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()

	r := SetupRotasTestes()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	urlDaBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", urlDaBusca, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAtualizaAluno(t *testing.T) {
	database.ConectaBanco()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupRotasTestes()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	urlEdicao := "/alunos/" + strconv.Itoa(ID)

	aluno := models.Aluno{Nome: "Aluno de teste", CPF: "1234567333", RG: "9876543333"}
	alunoJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", urlEdicao, bytes.NewBuffer(alunoJson))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Aluno de teste", alunoMock.Nome)
	assert.Equal(t, "1234567333", alunoMock.CPF)
	assert.Equal(t, "9876543333", alunoMock.RG)
	assert.Equal(t, http.StatusOK, res.Code)
}
