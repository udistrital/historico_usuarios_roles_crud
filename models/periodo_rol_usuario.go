package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/logs"
	"github.com/udistrital/utils_oas/time_bogota"
)

type PeriodoRolUsuario struct {
	Id                int      `orm:"column(id);pk;auto"`
	Activo            bool     `orm:"column(activo)"`
	FechaCreacion     string   `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string   `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	FechaInicio       string   `orm:"column(fecha_inicio);type(timestamp without time zone)"`
	FechaFin          *string  `orm:"column(fecha_fin);type(timestamp without time zone);null"`
	UsuarioId         *Usuario `orm:"column(usuario_id);rel(fk)"`
	RolId             *Rol     `orm:"column(rol_id);rel(fk)"`
}

func (t *PeriodoRolUsuario) TableName() string {
	return "periodo_rol_usuario"
}

func init() {
	orm.RegisterModel(new(PeriodoRolUsuario))
}

// AddPeriodoRolUsuario insert a new PeriodoRolUsuario into database and returns
// last inserted Id on success.
func AddPeriodoRolUsuario(m *PeriodoRolUsuario) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPeriodoRolUsuarioById retrieves PeriodoRolUsuario by Id. Returns error if
// Id doesn't exist
func GetPeriodoRolUsuarioById(id int) (v *PeriodoRolUsuario, err error) {
	o := orm.NewOrm()
	v = &PeriodoRolUsuario{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPeriodoRolUsuario retrieves all PeriodoRolUsuario matches certain condition. Returns empty list if
// no records exist
func GetAllPeriodoRolUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PeriodoRolUsuario)).RelatedSel()
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []PeriodoRolUsuario
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePeriodoRolUsuario updates PeriodoRolUsuario by Id and returns error if
// the record to be updated doesn't exist
func UpdatePeriodoRolUsuarioById(m *PeriodoRolUsuario) (err error) {
	o := orm.NewOrm()
	v := PeriodoRolUsuario{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePeriodoRolUsuario deletes PeriodoRolUsuario by Id and returns error if
// the record to be deleted doesn't exist
func DeletePeriodoRolUsuario(id int) (err error) {
	o := orm.NewOrm()
	v := PeriodoRolUsuario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {

		v.Activo = false
		v.FechaCreacion = time_bogota.TiempoCorreccionFormato(v.FechaCreacion)
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if _, err = o.Update(&v, "Activo", "FechaModificacion"); err == nil {
			fmt.Println("El registro ha sido marcado como inactivo")
		} else {
			fmt.Println("Error al actualizar el campo Activo:", err)
		}

	} else {
		fmt.Println("No exite el registro", err)
	}
	return
}
func ValidarAsignarPerido(idUsuario int, idRol int) error {
	var periodos []PeriodoRolUsuario
	o := orm.NewOrm()
	qs := o.QueryTable(new(PeriodoRolUsuario))
	num, err := qs.Filter("activo", true).Filter("usuario_id", idUsuario).Filter("rol_id", idRol).All(&periodos)
	if err == nil && num > 0 {
		logs.Info(num)
		return errors.New("el usuario ya tiene registrado el rol activo")
	}
	return nil
}
