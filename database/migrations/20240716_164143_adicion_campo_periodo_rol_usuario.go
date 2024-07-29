package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AdicionCampoPeriodoRolUsuario_20240716_164143 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AdicionCampoPeriodoRolUsuario_20240716_164143{}
	m.Created = "20240716_164143"

	migration.Register("AdicionCampoPeriodoRolUsuario_20240716_164143", m)
}

// Run the migrations
func (m *AdicionCampoPeriodoRolUsuario_20240716_164143) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20240716_164143_adicion_campo_periodo_rol_usuario.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *AdicionCampoPeriodoRolUsuario_20240716_164143) Down() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := os.ReadFile("../scripts/20240716_164143_adicion_campo_periodo_rol_usuario.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
