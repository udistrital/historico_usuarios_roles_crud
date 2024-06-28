package services

import (
	"testing"

	"github.com/udistrital/usuario_rol_crud/services"
)

func TestValidarTipoFecha(t *testing.T) {
	_, err := services.ValidarTipoFecha("2024-12-23 20:31:25.585123")
	if err == nil {
		t.Log("se esperaba valido ")
	} else {
		t.Error(" se esperaba error")
		t.Fail()
	}
}
