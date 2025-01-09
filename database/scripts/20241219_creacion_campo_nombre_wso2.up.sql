ALTER TABLE historico_rol_usuario.rol
ADD COLUMN nombre_wso2 VARCHAR(20);

COMMENT ON COLUMN historico_rol_usuario.rol.nombre_wso2 IS 'Nombre con que aparece registrado en WSO2';