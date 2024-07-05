package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/logs"
	"github.com/udistrital/usuario_rol_crud/helpers"
	"github.com/udistrital/usuario_rol_crud/models"
	"github.com/udistrital/usuario_rol_crud/services"

	"github.com/astaxie/beego"
)

// PeriodoRolUsuarioController operations for PeriodoRolUsuario
type PeriodoRolUsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *PeriodoRolUsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create PeriodoRolUsuario
// @Param	body		body 	models.PeriodoRolUsuario	true		"body for PeriodoRolUsuario content"
// @Success 201 {int} models.PeriodoRolUsuario
// @Failure 403 body is empty
// @router / [post]
func (c *PeriodoRolUsuarioController) Post() {
	defer helpers.ErrorController(c.Controller, "PeriodoRolUsuarioController")

	var v models.PeriodoRolUsuario
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if _, err := services.AddPeriodoRolUsuario(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 201, "Message": "Registro exitoso", "Data": v}

		} else {
			logs.Error(err)
			c.Data["Message"] = "Error servicio Post:  la petición contiene un parámetro incorrecto o no existe ningún registro"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error servicio Post:  la petición contiene un parámetro incorrecto o no existe ningún registro"
		c.Abort("400")
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get PeriodoRolUsuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PeriodoRolUsuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PeriodoRolUsuarioController) GetOne() {
	defer helpers.ErrorController(c.Controller, "PeriodoRolUsuarioController")

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := services.GetPeriodoRolUsuarioById(id)
	if err != nil {
		logs.Error(err)
		c.Data["Message"] = "Error en el servicio GetOne: la solicitud contiene un parámetro incorrecto o no existe ningún registro."
		c.Abort("400")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Petición exitosa", "Data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get PeriodoRolUsuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.PeriodoRolUsuario
// @Failure 403
// @router / [get]
func (c *PeriodoRolUsuarioController) GetAll() {
	defer helpers.ErrorController(c.Controller, "PeriodoRolUsuarioController")

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	//l, err := models.GetAllPeriodoRolUsuario(query, fields, sortby, order, offset, limit)
	l, err := services.GetAllPeriodoRolUsuario(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["Message"] = "Error servicio GetAll: la solicitud contiene un parámetro incorrecto o no existe ningún registro."
		c.Abort("404")
	} else {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Petición exitosa", "Data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the PeriodoRolUsuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.PeriodoRolUsuario	true		"body for PeriodoRolUsuario content"
// @Success 200 {object} models.PeriodoRolUsuario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PeriodoRolUsuarioController) Put() {
	defer helpers.ErrorController(c.Controller, "PeriodoRolUsuarioController")

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.PeriodoRolUsuario{Id: id}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := services.UpdatePeriodoRolUsuarioById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Modificacion exitosa", "Data": v}
		} else {
			logs.Error(err)
			c.Data["Message"] = "Error servicio Put: la solicitud contiene un tipo de datos incorrecto o un parámetro no válido"
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error servicio Put: la solicitud contiene un tipo de datos incorrecto o un parámetro no válido"
		c.Abort("400")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the PeriodoRolUsuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PeriodoRolUsuarioController) Delete() {
	defer helpers.ErrorController(c.Controller, "PeriodoRolUsuarioController")

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := services.DeletePeriodoRolUsuario(id); err == nil {
		d := map[string]interface{}{"Id": id}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": 200, "Message": "Eliminacion exitosa", "Data": d}
	} else {
		logs.Error(err)
		c.Data["Message"] = "Error servicio Delete: la solicitud contiene un tipo de datos incorrecto o un parámetro no válido"
		c.Abort("404")
	}
	c.ServeJSON()
}
