package services

import (
	"errors"
	"time"

	"github.com/udistrital/usuario_rol_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

type PeriodoRolUsuarioService struct{}

// crear nuevo periodo rol usuario
func AddPeriodoRolUsuario(v *models.PeriodoRolUsuario) (id int64, err error) {
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoBogotaFormato()
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()

	if v.FechaFin != nil {
		if err := validarPeriodoFechas(v.FechaInicio, v.FechaFin); err != nil {
			return 0, err
		}
	} else {
		_, err := validarTipoFecha(v.FechaInicio)
		v.FechaFin = nil
		if err != nil {
			return 0, err
		}
	}
	return models.AddPeriodoRolUsuario(v)
}
func GetPeriodoRolUsuarioById(id int) (v *models.PeriodoRolUsuario, err error) {
	return models.GetPeriodoRolUsuarioById(id)
}
func GetAllPeriodoRolUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return models.GetAllPeriodoRolUsuario(query, fields, sortby, order, offset, limit)
}
func UpdatePeriodoRolUsuarioById(v *models.PeriodoRolUsuario) (err error) {
	//se recupera periodo existente para mantener fecha de creacion
	periodo, err := models.GetPeriodoRolUsuarioById(v.Id)
	if err != nil {
		return err
	}
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoCorreccionFormato(periodo.FechaCreacion)
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()

	if v.FechaFin != nil {
		if err := validarPeriodoFechas(v.FechaInicio, v.FechaFin); err != nil {
			return err
		}
	} else {
		_, err := validarTipoFecha(v.FechaInicio)
		v.FechaFin = nil
		if err != nil {
			return err
		}
	}
	return models.UpdatePeriodoRolUsuarioById(v)
}
func DeletePeriodoRolUsuario(id int) (err error) {
	return models.DeletePeriodoRolUsuario(id)
}

// se valida el tipo de dato y estructura que ingresa para las fechas de inicio y fin
func validarTipoFecha(fecha string) (date time.Time, err error) {
	layout := "2006-01-02 15:04:05.999999"
	fechaParseada, err := time.Parse(layout, fecha)
	if err != nil {
		return fechaParseada, errors.New("formato de la fecha " + fecha + " es incorrecto")
	}
	return fechaParseada, nil
}

// se valida que la fecha de fin no vaya a ser menor que la de inicio
func validarPeriodoFechas(inicio string, fin *string) error {
	// Validar fecha de inicio
	inicioTime, err := validarTipoFecha(inicio)
	if err != nil {
		return err
	}

	// Validar fecha de fin
	fechaFin := *fin
	finTime, err := validarTipoFecha(fechaFin)
	if err != nil {
		return err
	}

	// Validar que la fecha de fin no sea anterior a la fecha de inicio
	if finTime.Before(inicioTime) {
		return errors.New("la fecha de fin no puede ser menor que la fecha de inicio")
	}
	return nil
}
