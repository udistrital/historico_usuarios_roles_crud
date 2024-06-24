package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:PeriodoRolUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:RolController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:SistemaInformacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/usuario_rol_crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
