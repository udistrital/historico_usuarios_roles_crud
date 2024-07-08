package controller

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostOk(t *testing.T) {
	jsonStr := []byte(`{"Documento":"77788891"}`)

	a := assert.New(t)
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/usuarios", bytes.NewBuffer(jsonStr))

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
	a.Equal(201, response.StatusCode, "correcto")
}
func TestPostEmpty(t *testing.T) {
	jsonStr := []byte(`{"Documento":""}`)

	a := assert.New(t)
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/usuarios", bytes.NewBuffer(jsonStr))

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
func TestPostFail(t *testing.T) {
	jsonStr := []byte(`{"Documento":"asd"}`)

	a := assert.New(t)
	req, err := http.NewRequest("POST", "http://localhost:8080/v1/usuarios", bytes.NewBuffer(jsonStr))

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
