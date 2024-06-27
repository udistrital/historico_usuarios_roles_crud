-- SCHEMA: usuario-rol

-- DROP SCHEMA IF EXISTS "usuario-rol" ;

CREATE SCHEMA IF NOT EXISTS "usuario_rol"
    AUTHORIZATION postgres;
	
	
	CREATE TABLE IF NOT EXISTS usuario_rol.usuario(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		documento varchar(20),
		CONSTRAINT pk_usuario PRIMARY KEY (id)
	);
	COMMENT ON TABLE  usuario_rol.usuario IS 'Tabla que almacena el documento del usuario, la informacion adicional sera traida de servicio terceros';
	COMMENT ON COLUMN  usuario_rol.usuario.id IS 'Identificador de la tabla de usuario';
	COMMENT ON COLUMN  usuario_rol.usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol.usuario.fecha_creacion IS 'Fecha de creacion de un usuario';
	COMMENT ON COLUMN  usuario_rol.usuario.fecha_modificacion IS 'Fecha de modifiación de un usuario';
	COMMENT ON COLUMN  usuario_rol.usuario.documento IS 'Campo para almacenar el documento de identidad del usuario';
	
	CREATE TABLE IF NOT EXISTS usuario_rol.sistema_informacion(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		nombre varchar(25) NOT NULL,
		descripcion varchar(255),
		CONSTRAINT pk_sistema_informacion PRIMARY KEY (id)
	);
	COMMENT ON TABLE  usuario_rol.sistema_informacion IS 'Tabla que almacena los sistemas de informacion';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.id IS 'Identificador de la tabla de roles';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.fecha_creacion IS 'Fecha de creacion de un rol';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.fecha_modificacion IS 'Fecha de modifiación de un rol';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.nombre IS 'Nombre del sistema de informacion';
	COMMENT ON COLUMN  usuario_rol.sistema_informacion.descripcion IS 'Descripcion del sistema de informacion';

	
	CREATE TABLE IF NOT EXISTS usuario_rol.rol(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		nombre varchar(20) NOT NULL,
		sistema_informacion_id integer NOT NULL,
		CONSTRAINT pk_rol PRIMARY KEY (id),
		CONSTRAINT fk_rol_sistema_informacion FOREIGN KEY (sistema_informacion_id) REFERENCES usuario_rol.sistema_informacion(id)
	);
	COMMENT ON TABLE  usuario_rol.rol IS 'Tabla que almacena los roles';
	COMMENT ON COLUMN  usuario_rol.rol.id IS 'Identificador de la tabla de roles';
	COMMENT ON COLUMN  usuario_rol.rol.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol.rol.fecha_creacion IS 'Fecha de creacion de un rol';
	COMMENT ON COLUMN  usuario_rol.rol.fecha_modificacion IS 'Fecha de modifiación de un rol';
	COMMENT ON COLUMN  usuario_rol.rol.nombre IS 'Nombre del rol';
	COMMENT ON COLUMN  usuario_rol.rol.sistema_informacion_id IS 'Id del sistema de informacion al que pertenece';
	
	
	CREATE TABLE IF NOT EXISTS usuario_rol.periodo_rol_usuario(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		fecha_inicio timestamp NOT NULL,
		fecha_fin timestamp,
		usuario_id integer NOT NULL,
		rol_id integer NOT NULL,
		CONSTRAINT pk_periodo_rol_usuario PRIMARY KEY (id),
		CONSTRAINT fk_periodo_rol_usuarios FOREIGN KEY (usuario_id) REFERENCES usuario_rol.usuario(id),
		CONSTRAINT fk_periodo_rol_roles FOREIGN KEY (rol_id) REFERENCES usuario_rol.rol(id)
	);
	COMMENT ON TABLE  usuario_rol.periodo_rol_usuario IS 'Tabla que almacena el estado y periodo de asignacion de los roles a los usuarios';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.id IS 'Identificador de la tabla';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.fecha_creacion IS 'Fecha de creacion del registro';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.fecha_modificacion IS 'Fecha de modifiación del registro';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.fecha_inicio IS 'Fecha en la que se le fue asignado el rol al usuario';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.fecha_fin IS 'Fecha en la que se le es designado el rol';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.usuario_id IS 'Usuario al que hace referencia';
	COMMENT ON COLUMN  usuario_rol.periodo_rol_usuario.rol_id IS 'ID del rol al que hace referencia';