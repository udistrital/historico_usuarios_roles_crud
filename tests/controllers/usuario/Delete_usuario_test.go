package controller

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteOk(t *testing.T) {
	a := assert.New(t)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/usuarios/10", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Comprobar el código de estado
	a.Equal(200, response.StatusCode, "correcto")
}
func TestDeleteNotFoud(t *testing.T) {
	a := assert.New(t)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/usuarios/88", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Comprobar el código de estado
	a.Equal(400, response.StatusCode, "correcto")
}

func TestDeleteFail(t *testing.T) {
	a := assert.New(t)
	req, err := http.NewRequest("DELETE", "http://localhost:8080/v1/usuarios/asd", nil)

	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	// Comprobar el código de estado
	a.Equal(400, response.StatusCode, "correcto")
}
