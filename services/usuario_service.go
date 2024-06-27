package services

import (
	"errors"

	"github.com/udistrital/usuario_rol_crud/models"
	"github.com/udistrital/utils_oas/time_bogota"
)

type Usuario struct{}

func AddUsuario(v *models.Usuario) (id int64, err error) {
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoBogotaFormato()
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	if models.DocumentoExistente(v.Documento) {
		return 0, errors.New("el documento que intenta registrar ya existe")
	}
	return models.AddUsuario(v)
}
func GetUsuarioById(id int) (v *models.Usuario, err error) {
	return models.GetUsuarioById(id)
}
func GetAllUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return models.GetAllUsuario(query, fields, sortby, order, offset, limit)
}
func UpdateUsuarioById(v *models.Usuario) (err error) {
	//se recupera usuario existente para mantener fecha de creacion
	periodo, err := models.GetUsuarioById(v.Id)
	if err != nil {
		return err
	}
	v.Activo = true
	v.FechaCreacion = time_bogota.TiempoCorreccionFormato(periodo.FechaCreacion)
	v.FechaModificacion = time_bogota.TiempoBogotaFormato()
	return models.UpdateUsuarioById(v)
}
func DeleteUsuario(id int) (err error) {
	return models.DeleteUsuario(id)
}
