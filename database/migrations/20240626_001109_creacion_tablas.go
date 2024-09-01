package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreacionTablas_20240626_001109 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreacionTablas_20240626_001109{}
	m.Created = "20240626_001109"

	migration.Register("CreacionTablas_20240626_001109", m)
}

// Run the migrations
func (m *CreacionTablas_20240626_001109) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	//se remueve ya que las migraciones se haran de manera manual y asi evitar conflictos
	/*file, err := os.ReadFile("../scripts/20240626_001109_creacion_tablas.up.sql")

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
func (m *CreacionTablas_20240626_001109) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	//se remueve ya que las migraciones se haran de manera manual y asi evitar conflictos

	/*file, err := os.ReadFile("../scripts/20240626_001109_creacion_tablas.down.sql")

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
