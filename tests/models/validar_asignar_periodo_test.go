package services

import (
	"testing"

	"github.com/udistrital/usuario_rol_crud/models"
)

func TestValidarAsignarPerido(t *testing.T) {
	err := models.ValidarAsignarPerido(1, 1)
	if err == nil {
		t.Log("se esperaba valido ")
	} else {
		t.Error("fallo")
		t.Fail()
	}
}
