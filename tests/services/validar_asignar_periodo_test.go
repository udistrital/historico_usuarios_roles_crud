package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/udistrital/usuario_rol_crud/services"
)

func TestValidarTipoFechaOk(t *testing.T) {
	a := assert.New(t)
	//_, err := services.ValidarTipoFecha("2024-12-23 20:31:25.585123")
	_, err := services.ValidarTipoFecha("2024-12-23")
	a.Equal(nil, err, "tipo de fecha correcta")
}

func TestValidarTipoFechaNoValido(t *testing.T) {
	a := assert.New(t)
	_, err := services.ValidarTipoFecha("asd123")
	a.Error(err, "se esperaba error")
	a.EqualError(err, "formato de la fecha asd123 es incorrecto", "valor incorrecto controlado")

}
func TestValidarTipoFechaIncompleto(t *testing.T) {
	a := assert.New(t)
	_, err := services.ValidarTipoFecha("2024-12-23")
	a.Error(err, "se esperaba error")
	a.EqualError(err, "formato de la fecha 2024-12-23 es incorrecto", "formato de fecha incorrecto controlado")

}
func TestValidarTipoFechaVacio(t *testing.T) {
	a := assert.New(t)
	_, err := services.ValidarTipoFecha("")
	a.Error(err, "se esperaba error")
	a.EqualError(err, "formato de la fecha  es incorrecto", "parametro de fecha vacio controlado")

}
