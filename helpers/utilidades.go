package helpers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
)

func ErrorController(c beego.Controller, controller string) {
	if err := recover(); err != nil {
		logs.Error(err)
		var localError map[string]interface{}
		switch e := err.(type) {
		case string:
			localError = map[string]interface{}{
				"funcion": "unknown",
				"err":     e,
				"status":  "500",
			}
		case error:
			localError = map[string]interface{}{
				"funcion": "unknown",
				"err":     e.Error(),
				"status":  "500",
			}
		default:
			if le, ok := err.(map[string]interface{}); ok {
				localError = le
			} else {
				localError = map[string]interface{}{
					"funcion": "unknown",
					"err":     "Unknown error",
					"status":  "500",
				}
			}
		}
		c.Data["message"] = beego.AppConfig.String("appname") + "/" + controller + "/" + localError["funcion"].(string)
		c.Data["data"] = localError["err"]
		if status, ok := localError["status"]; ok {
			c.Ctx.Output.SetStatus(parseStatus(status.(string))) // Establecer el código de estado HTTP correctamente
			c.ServeJSON()
		} else {
			c.Abort("500")
		}
	}
}

func parseStatus(status string) int {
	parsedStatus, err := strconv.Atoi(status)
	if err != nil {
		return 500
	}
	return parsedStatus
}
