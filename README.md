# historico_usuarios_roles
El APIes un apoyo para el servicio de WSO2 gestiona los diferentes usuarios y sus respectivos roles en cada sistema de informacion 

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang 1.22.2](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo 1.12.0](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)

### Variables de Entorno

usuario_rol_v1_PGuser=[usuario con acceso a la base de datos]
usuario_rol_v1_PGpass=[password del usuario]
usuario_rol_v1_PGhost=[Puerto de conexión con la base de datos]
usuario_rol_v1_PGport=[puerto de ejecucion]
usuario_rol_v1_PGdb=[nombre de la base de datos]
usuario_rol_v1_PGschema=[esquema donde se ubican las tablas]

### Ejecución del Proyecto
shel

# 1. Obtener el repositorio con Go
go get github.com/udistrital/inscripcion_crud

# 2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/inscripcion_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.

# 5. ejecutar el proyecto
bee run -downdoc=true -gendoc=true 

## Modelo de Datos

[Modelo de Datos API CRUD historico_usuarios_roles_crud](https://udistritaleduco-my.sharepoint.com/:i:/g/personal/computo_udistrital_edu_co/EVcG7HT2O4ZJg-XsKu4XsaYBqSijATtERNxWIXvxRhfJZw?e=f2BdIN)

