package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambioTipodatoFechasInicioFin_20240726_155105 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambioTipodatoFechasInicioFin_20240726_155105{}
	m.Created = "20240726_155105"

	migration.Register("CambioTipodatoFechasInicioFin_20240726_155105", m)
}

// Run the migrations
func (m *CambioTipodatoFechasInicioFin_20240726_155105) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	//se remueve ya que las migraciones se haran de manera manual y asi evitar conflictos
	/*file, err := os.ReadFile("../scripts/20240726_155105_cambio_tipodato_fechas_inicio_fin.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}*/
}

// Reverse the migrations
func (m *CambioTipodatoFechasInicioFin_20240726_155105) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	//se remueve ya que las migraciones se haran de manera manual y asi evitar conflictos
	/*file, err := os.ReadFile("../scripts/20240726_155105_cambio_tipodato_fechas_inicio_fin.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}*/
}
