ALTER TABLE periodo_rol_usuario ALTER COLUMN fecha_inicio TYPE TIMESTAMP USING fecha_inicio::TIMESTAMP;
ALTER TABLE periodo_rol_usuario ALTER COLUMN fecha_fin TYPE TIMESTAMP USING fecha_fin::TIMESTAMP;
