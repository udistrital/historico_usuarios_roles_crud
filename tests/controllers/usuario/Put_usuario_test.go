package controller

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPutOk(t *testing.T) {
	jsonStr := []byte(`{"Documento":"77788891"}`)

	a := assert.New(t)
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/usuarios/1", bytes.NewBuffer(jsonStr))

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
func TestPutNotFoud(t *testing.T) {
	jsonStr := []byte(`{"Documento":"77788891"}`)

	a := assert.New(t)
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/usuarios/88", bytes.NewBuffer(jsonStr))

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

func TestPutFail(t *testing.T) {
	jsonStr := []byte(`{"Documento":"77788891"}`)

	a := assert.New(t)
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/usuarios/asd", bytes.NewBuffer(jsonStr))

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
func TestPutJsonError(t *testing.T) {
	jsonStr := []byte(`{"Documento":"asdc"}`)

	a := assert.New(t)
	req, err := http.NewRequest("PUT", "http://localhost:8080/v1/usuarios/1", bytes.NewBuffer(jsonStr))

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
