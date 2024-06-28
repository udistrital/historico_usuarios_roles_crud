package services

import (
	"testing"

	"github.com/udistrital/usuario_rol_crud/services"
)

func TestValidarPeriodoFechas(t *testing.T) {
	var1 := "2024-12-23 20:31:25.585123"
	var2 := "2025-12-23 20:31:25.585123"
	err := services.ValidarPeriodoFechas(var1, &var2)
	if err == nil {
		t.Log("se esperaba valido ")
	} else {
		t.Error(" se esperaba error")
		t.Fail()
	}
}
