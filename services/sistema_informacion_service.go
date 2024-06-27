package services

import (
	"github.com/udistrital/usuario_rol_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

type SistemaInformacion struct{}

func AddSistemaInformacion(v *models.SistemaInformacion) (id int64, err error) {
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoBogotaFormato()
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	return models.AddSistemaInformacion(v)
}
func GetSistemaInformacionById(id int) (v *models.SistemaInformacion, err error) {
	return models.GetSistemaInformacionById(id)
}
func GetAllSistemaInformacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return models.GetAllSistemaInformacion(query, fields, sortby, order, offset, limit)
}
func UpdateSistemaInformacionById(v *models.SistemaInformacion) (err error) {
	periodo, err := models.GetSistemaInformacionById(v.Id)
	if err != nil {
		return err
	}
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoCorreccionFormato(periodo.FechaCreacion)
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	return models.UpdateSistemaInformacionById(v)
}
func DeleteSistemaInformacion(id int) (err error) {
	return models.DeleteSistemaInformacion(id)

}
