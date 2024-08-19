package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type PeriodoRolUsuario struct {
	Id                int      `orm:"column(id);pk;auto"`
	Activo            bool     `orm:"column(activo)"`
	FechaCreacion     string   `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string   `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	FechaInicio       string   `orm:"column(fecha_inicio);type(date)"`
	FechaFin          *string  `orm:"column(fecha_fin);type(date);null"`
	Finalizado        bool     `orm:"column(finalizado)"`
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
func GetPeriodoRolUsuarioById(id int) (*PeriodoRolUsuario, error) {
	o := orm.NewOrm()
	v := &PeriodoRolUsuario{Id: id}

	// Leer la entidad principal
	if err := o.Read(v); err != nil {
		return nil, err
	}

	// Cargar Usuario relacionado
	if v.UsuarioId != nil {
		if err := o.Read(v.UsuarioId); err != nil {
			return nil, err
		}
	}

	// Cargar Rol relacionado
	if v.RolId != nil {
		if err := o.Read(v.RolId); err != nil {
			return nil, err
		}
	}

	// cargar el SistemaInformacion relacionado con Rol
	if v.RolId != nil && v.RolId.SistemaInformacionId != nil {
		if err := o.Read(v.RolId.SistemaInformacionId); err != nil {
			return nil, err
		}
	}

	// Formatear fechas
	formateado, err := formatoFechas(*v)
	if err != nil {
		return nil, err
	}

	return &formateado, nil
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
				formateado, err := formatoFechas(v)
				if err != nil {
					return nil, err
				}
				ml = append(ml, formateado)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				formateado, err := formatoFechas(v)
				if err != nil {
					return nil, err
				}
				m := make(map[string]interface{})
				val := reflect.ValueOf(formateado)
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
func formatoFechas(p PeriodoRolUsuario) (PeriodoRolUsuario, error) {
	const layoutDateTime = "2006-01-02 15:04:05 +0000 +0000"

	if p.FechaInicio != "" {
		t, err := time.Parse(layoutDateTime, p.FechaInicio)
		if err != nil {
			return p, err
		}
		p.FechaInicio = t.Format("2006-01-02") // Solo la fecha
	}

	if p.FechaFin != nil && *p.FechaFin != "" {
		t, err := time.Parse(layoutDateTime, *p.FechaFin)
		if err != nil {
			return p, err
		}
		formatted := t.Format("2006-01-02") // Solo la fecha
		p.FechaFin = &formatted
	}
	return p, nil
}

// GetPeriodosByUsuarioId obtiene los periodos de un usuario dado su ID
func GetPeriodosByUsuarioId(usuarioId int, query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {

	o := orm.NewOrm()
	//var periodos []PeriodoRolUsuario
	//_, err := o.QueryTable("periodo_rol_usuario").Filter("usuario_id", usuarioId).All(&periodos)
	qs := o.QueryTable(new(PeriodoRolUsuario)).RelatedSel().Filter("usuario_id", usuarioId)
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
				formateado, err := formatoFechas(v)
				if err != nil {
					return nil, err
				}
				ml = append(ml, formateado)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				formateado, err := formatoFechas(v)
				if err != nil {
					return nil, err
				}
				m := make(map[string]interface{})
				val := reflect.ValueOf(formateado)
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
func GetPeriodosBySistemaId(usuarioId *string, sistemaId string, query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	var results []PeriodoRolUsuario
	sql := `SELECT p.* FROM usuario_rol.periodo_rol_usuario p 
		INNER JOIN usuario_rol.rol r ON p.rol_id = r.id 
		INNER JOIN usuario_rol.sistema_informacion s ON r.sistema_informacion_id = s.id 
		INNER JOIN usuario_rol.usuario u ON p.usuario_id = u.id
		WHERE `

	sql += fmt.Sprintf("s.id = '%s'", sistemaId)

	// Si se pasa un usuarioId, se agrega a la consulta
	if usuarioId != nil {
		sql += fmt.Sprintf(" AND u.documento = '%s'", *usuarioId)
	}

	// Agregar otros filtros dinámicos
	for k, v := range query {
		// Omitir agregar "sistema_informacion" porque ya está siendo usado en "s.id = ?"
		if k == "sistema_informacion" {
			continue
		}
		sql += fmt.Sprintf(" AND %s = '%s'", k, v)
	}

	// Sorting
	if len(sortby) > 0 {
		sql += " ORDER BY "
		for i, s := range sortby {
			orderby := s
			if len(order) == 1 {
				orderby += " " + order[0]
			} else if len(order) > 1 {
				orderby += " " + order[i]
			}
			if i < len(sortby)-1 {
				sql += orderby + ", "
			} else {
				sql += orderby
			}
		}
	}

	// Pagination
	sql += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	_, err = o.Raw(sql).QueryRows(&results)
	if err != nil {
		return nil, err
	}

	// Llenar las relaciones
	for i := range results {
		o.LoadRelated(&results[i], "UsuarioId")
		o.LoadRelated(&results[i], "RolId")
		if results[i].RolId != nil {
			o.LoadRelated(results[i].RolId, "SistemaInformacionId")
		}
	}

	// Convertir los resultados a []interface{}
	for _, v := range results {
		if len(fields) > 0 {
			m := make(map[string]interface{})
			val := reflect.ValueOf(v)
			for _, fname := range fields {
				m[fname] = val.FieldByName(fname).Interface()
			}
			ml = append(ml, m)
		} else {
			ml = append(ml, v)
		}
	}
	return ml, nil
}
