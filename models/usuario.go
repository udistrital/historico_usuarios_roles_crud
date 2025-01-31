package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/udistrital/utils_oas/time_bogota"
)

type Usuario struct {
	Id                int    `orm:"column(id);pk;auto"`
	Activo            bool   `orm:"column(activo)"`
	FechaCreacion     string `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion string `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	Documento         string `orm:"column(documento);null"`
}

func (t *Usuario) TableName() string {
	return "usuario"
}

func init() {
	orm.RegisterModel(new(Usuario))
}

// AddUsuario insert a new Usuario into database and returns
// last inserted Id on success.
func AddUsuario(m *Usuario) (id int64, err error) {
	o := orm.NewOrm()
	//se valida si el numero de documento ya existe antes de registrarlo
	exist := Usuario{Documento: m.Documento}
	if m.Documento != "" && o.Read(&exist, "Documento") == nil {
		return 0, errors.New("documento already exists")
	}
	id, err = o.Insert(m)
	return
}

// GetUsuarioById retrieves Usuario by Id. Returns error if
// Id doesn't exist
func GetUsuarioById(id int) (v *Usuario, err error) {
	o := orm.NewOrm()
	v = &Usuario{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsuario retrieves all Usuario matches certain condition. Returns empty list if
// no records exist
func GetAllUsuario(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Usuario)).RelatedSel()
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

	var l []Usuario
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

// UpdateUsuario updates Usuario by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsuarioById(m *Usuario) (err error) {
	o := orm.NewOrm()
	v := Usuario{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsuario deletes Usuario by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsuario(id int) (err error) {
	o := orm.NewOrm()
	v := Usuario{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {

		v.Activo = false
		v.FechaCreacion = time_bogota.TiempoCorreccionFormato(v.FechaCreacion)
		v.FechaModificacion = time_bogota.TiempoBogotaFormato()
		if _, err = o.Update(&v, "activo", "fechaModificacion"); err == nil {
			fmt.Println("El registro ha sido marcado como inactivo")
		} else {
			fmt.Println("Error al actualizar el campo Activo:", err)
		}

	} else {
		logs.Error(err)
		fmt.Println("No exite el registro", err)
	}
	return
}

// Check if a Documento already exists
func DocumentoExistente(documento string) bool {
	o := orm.NewOrm()
	return o.QueryTable("usuario").Filter("documento", documento).Exist()
}
func GetUsuarioByDocumento(documento string) (*Usuario, error) {
	o := orm.NewOrm()
	var usuario Usuario
	err := o.QueryTable("usuario").Filter("documento", documento).One(&usuario)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
