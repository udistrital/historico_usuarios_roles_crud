-- SCHEMA: historico_rol_usuario

DROP SCHEMA IF EXISTS "historico_rol_usuario" ;

CREATE SCHEMA IF NOT EXISTS "historico_rol_usuario"
  	
	CREATE TABLE IF NOT EXISTS historico_rol_usuario.usuario(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		documento varchar(20),
		CONSTRAINT pk_usuario PRIMARY KEY (id)
	);
	COMMENT ON TABLE  historico_rol_usuario.usuario IS 'Tabla que almacena el documento del usuario, la informacion adicional sera traida de servicio terceros';
	COMMENT ON COLUMN  historico_rol_usuario.usuario.id IS 'Identificador de la tabla de usuario';
	COMMENT ON COLUMN  historico_rol_usuario.usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  historico_rol_usuario.usuario.fecha_creacion IS 'Fecha de creacion de un usuario';
	COMMENT ON COLUMN  historico_rol_usuario.usuario.fecha_modificacion IS 'Fecha de modifiación de un usuario';
	COMMENT ON COLUMN  historico_rol_usuario.usuario.documento IS 'Campo para almacenar el documento de identidad del usuario';
	
	CREATE TABLE IF NOT EXISTS historico_rol_usuario.rol(
		id serial NOT NULL,
		activo boolean NOT NULL,
		fecha_creacion timestamp NOT NULL,
		fecha_modificacion timestamp NOT NULL,
		nombre varchar(20) NOT NULL,
		sistema_informacion_id integer NOT NULL,
		CONSTRAINT pk_rol PRIMARY KEY (id)
	);
	COMMENT ON TABLE  historico_rol_usuario.rol IS 'Tabla que almacena los roles';
	COMMENT ON COLUMN  historico_rol_usuario.rol.id IS 'Identificador de la tabla de roles';
	COMMENT ON COLUMN  historico_rol_usuario.rol.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  historico_rol_usuario.rol.fecha_creacion IS 'Fecha de creacion de un rol';
	COMMENT ON COLUMN  historico_rol_usuario.rol.fecha_modificacion IS 'Fecha de modifiación de un rol';
	COMMENT ON COLUMN  historico_rol_usuario.rol.nombre IS 'Nombre del rol';
	COMMENT ON COLUMN  historico_rol_usuario.rol.sistema_informacion_id IS 'Id del sistema de informacion al que pertenece';
	
	
	CREATE TABLE IF NOT EXISTS historico_rol_usuario.periodo_rol_usuario(
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
		CONSTRAINT fk_periodo_rol_usuarios FOREIGN KEY (usuario_id) REFERENCES historico_rol_usuario.usuario(id),
		CONSTRAINT fk_periodo_rol_roles FOREIGN KEY (rol_id) REFERENCES historico_rol_usuario.rol(id)
	);
	COMMENT ON TABLE  historico_rol_usuario.periodo_rol_usuario IS 'Tabla que almacena el estado y periodo de asignacion de los roles a los usuarios';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.id IS 'Identificador de la tabla';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.activo IS 'Campo para identificar si el  registro se encuentra activo o no, solo a nivel de registro.';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.fecha_creacion IS 'Fecha de creacion del registro';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.fecha_modificacion IS 'Fecha de modifiación del registro';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.fecha_inicio IS 'Fecha en la que se le fue asignado el rol al usuario';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.fecha_fin IS 'Fecha en la que se le es designado el rol';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.usuario_id IS 'Usuario al que hace referencia';
	COMMENT ON COLUMN  historico_rol_usuario.periodo_rol_usuario.rol_id IS 'ID del rol al que hace referencia';