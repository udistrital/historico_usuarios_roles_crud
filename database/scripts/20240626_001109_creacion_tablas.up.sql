-- SCHEMA: usuario_rol_db

DROP SCHEMA IF EXISTS "usuario_rol_db" ;

CREATE SCHEMA IF NOT EXISTS "usuario_rol_db"
  	
	CREATE TABLE IF NOT EXISTS usuario_rol_db.usuario(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		documento varchar(20),
		CONSTRAINT pk_usuario PRIMARY KEY (id)
	);
	COMMENT ON TABLE  usuario_rol_db.usuario IS 'Tabla que almacena el documento del usuario, la informacion adicional sera traida de servicio terceros';
	COMMENT ON COLUMN  usuario_rol_db.usuario.id IS 'Identificador de la tabla de usuario';
	COMMENT ON COLUMN  usuario_rol_db.usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol_db.usuario.fecha_creacion IS 'Fecha de creacion de un usuario';
	COMMENT ON COLUMN  usuario_rol_db.usuario.fecha_modificacion IS 'Fecha de modifiación de un usuario';
	COMMENT ON COLUMN  usuario_rol_db.usuario.documento IS 'Campo para almacenar el documento de identidad del usuario';
	
	CREATE TABLE IF NOT EXISTS usuario_rol_db.rol(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		nombre varchar(20) NOT NULL,
		sistema_informacion_id integer NOT NULL,
		CONSTRAINT pk_rol PRIMARY KEY (id)
	);
	COMMENT ON TABLE  usuario_rol_db.rol IS 'Tabla que almacena los roles';
	COMMENT ON COLUMN  usuario_rol_db.rol.id IS 'Identificador de la tabla de roles';
	COMMENT ON COLUMN  usuario_rol_db.rol.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol_db.rol.fecha_creacion IS 'Fecha de creacion de un rol';
	COMMENT ON COLUMN  usuario_rol_db.rol.fecha_modificacion IS 'Fecha de modifiación de un rol';
	COMMENT ON COLUMN  usuario_rol_db.rol.nombre IS 'Nombre del rol';
	COMMENT ON COLUMN  usuario_rol_db.rol.sistema_informacion_id IS 'Id del sistema de informacion al que pertenece';
	
	
	CREATE TABLE IF NOT EXISTS usuario_rol_db.periodo_rol_usuario(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		fecha_inicio date NOT NULL,
		fecha_fin date,
		usuario_id integer NOT NULL,
		rol_id integer NOT NULL,
		finalizado BOOLEAN NOT NULL,
		CONSTRAINT pk_periodo_rol_usuario PRIMARY KEY (id),
		CONSTRAINT fk_periodo_rol_usuarios FOREIGN KEY (usuario_id) REFERENCES usuario_rol_db.usuario(id),
		CONSTRAINT fk_periodo_rol_roles FOREIGN KEY (rol_id) REFERENCES usuario_rol_db.rol(id)
	);
	COMMENT ON TABLE  usuario_rol_db.periodo_rol_usuario IS 'Tabla que almacena el estado y periodo de asignacion de los roles a los usuarios';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.id IS 'Identificador de la tabla';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.fecha_creacion IS 'Fecha de creacion del registro';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.fecha_modificacion IS 'Fecha de modifiación del registro';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.fecha_inicio IS 'Fecha en la que se le fue asignado el rol al usuario';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.fecha_fin IS 'Fecha en la que se le es designado el rol';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.usuario_id IS 'Usuario al que hace referencia';
	COMMENT ON COLUMN  usuario_rol_db.periodo_rol_usuario.rol_id IS 'ID del rol al que hace referencia';