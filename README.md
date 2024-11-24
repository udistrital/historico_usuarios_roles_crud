# Usuario Rol
El API permite gestionar el histórico de roles de un usuario almacenando el periodo, rol y el id del usuario por sistema de información.
Puede ser usado como apoyo para el servicio de WSO2 garantizando el almacenamiento de los cambios en roles efectuados.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang 1.22.2](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo 1.12.0](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

USUARIO_ROL_PGuser=[usuario con acceso a la base de datos]
USUARIO_ROL_PGpass=[password del usuario]
USUARIO_ROL_PGhost=[Puerto de conexión con la base de datos]
USUARIO_ROL_PGport=[puerto de ejecucion]
USUARIO_ROL_PGdb=[nombre de la base de datos]
USUARIO_ROL_PGschema=[esquema donde se ubican las tablas]

### Ejecución del Proyecto
```shel

# 1. Obtener el repositorio con Go
go get github.com/udistrital/inscripcion_crud

# 2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/inscripcion_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.

# 5. ejecutar el proyecto
bee run -downdoc=true -gendoc=true 
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/usuario_rol_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/usuario_rol_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/usuario_rol_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/usuario_rol_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/usuario_rol_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/usuario_rol_crud/) |

## Modelo de Datos

[Modelo de Datos API CRUD usuario_rol_crud](./docs/usuarios%20y%20roles%20V6.drawio.png)


## Licencia

This file is part of usuario_rol_crud.

usuario_rol_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

usuario_rol_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.
