ALTER TABLE periodo_rol_usuario ALTER COLUMN fecha_inicio TYPE DATE USING fecha_inicio::DATE;
ALTER TABLE periodo_rol_usuario ALTER COLUMN fecha_fin TYPE DATE USING fecha_fin::DATE;