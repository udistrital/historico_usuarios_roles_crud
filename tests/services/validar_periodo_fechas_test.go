package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udistrital/usuario_rol_crud/services"
)

func TestValidarPeriodoFechasOK(t *testing.T) {
	a := assert.New(t)
	var1 := "2024-11-23"
	var2 := "2025-12-23"
	err := services.ValidarPeriodoFechas(var1, &var2)
	a.Equal(nil, err, "periodo de fechas correcto Controlado")
}

/*func TestValidarPeriodoFechasFallo(t *testing.T) {
	a := assert.New(t)
	var1 := "2024-12-23"
	var2 := "2022-12-23"
	err := services.ValidarPeriodoFechas(var1, &var2)
	assert.NotNil(t, err)

	expectedMessage := "la fecha de fin no puede ser menor que la fecha de inicio"

	a.Equal(t, err, expectedMessage)
}*/
