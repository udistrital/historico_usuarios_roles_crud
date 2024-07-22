# historico_usuarios_roles
El APIes un apoyo para el servicio de WSO2 gestiona los diferentes usuarios y sus respectivos roles en cada sistema de informacion 

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang 1.22.2](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo 1.12.0](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno

usuario_rol_v1_PGuser=[usuario con acceso a la base de datos]
usuario_rol_v1_PGpass=[password del usuario]
usuario_rol_v1_PGhost=[Puerto de conexión con la base de datos]
usuario_rol_v1_PGport=[puerto de ejecucion]
usuario_rol_v1_PGdb=[nombre de la base de datos]
usuario_rol_v1_PGschema=[esquema donde se ubican las tablas]

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
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/historico_usuarios_roles_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/historico_usuarios_roles_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/historico_usuarios_roles_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/historico_usuarios_roles_crud/) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/historico_usuarios_roles_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/historico_usuarios_roles_crud/) |

## Modelo de Datos

[Modelo de Datos API CRUD historico_usuarios_roles_crud](https://udistritaleduco-my.sharepoint.com/:i:/g/personal/computo_udistrital_edu_co/ETC3Qy30I-xMqof1NbjVrikBXokuoCPu-HuAf4Spz34l3w?e=9bmgpi)


## Licencia

This file is part of historico_usuarios_roles_crud.

inscripcion_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

inscripcion_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with novedades_crud. If not, see https://www.gnu.org/licenses/.