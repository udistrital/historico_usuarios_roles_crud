package controller

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOneOk(t *testing.T) {
	a := assert.New(t)
	response, _ := http.Get("http://localhost:8080/v1/usuarios/1")
	a.Equal(200, response.StatusCode, "correcto")
}
func TestGetOneFail(t *testing.T) {
	a := assert.New(t)
	response, _ := http.Get("http://localhost:8080/v1/usuarios/asd")
	a.Equal(404, response.StatusCode, "correcto")
}
func TestGetOneNotFound(t *testing.T) {
	a := assert.New(t)
	response, _ := http.Get("http://localhost:8080/v1/usuarios/88")
	a.Equal(404, response.StatusCode, "correcto")
}
