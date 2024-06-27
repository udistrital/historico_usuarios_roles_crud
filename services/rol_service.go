package services

import (
	"github.com/udistrital/usuario_rol_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

type Rol struct{}

func AddRol(v *models.Rol) (id int64, err error) {
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoBogotaFormato()
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	return models.AddRol(v)
}
func GetRolById(id int) (v *models.Rol, err error) {
	return models.GetRolById(id)
}

func GetAllRol(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return models.GetAllRol(query, fields, sortby, order, offset, limit)
}

func UpdateRolById(v *models.Rol) (err error) {
	//se recupera usuario existente para mantener fecha de creacion
	periodo, err := models.GetRolById(v.Id)
	if err != nil {
		return err
	}
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoCorreccionFormato(periodo.FechaCreacion)
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	return models.UpdateRolById(v)
}

func DeleteRol(id int) (err error) {
	return models.DeleteRol(id)
}
