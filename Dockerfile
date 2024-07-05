# Etapa 1: Compilación
FROM golang:1.22 as builder
 
# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app
 
# Copiar los archivos del proyecto al directorio de trabajo
COPY . .
 
# Descargar las dependencias del proyecto
RUN go mod download
 
# Compilar la aplicación Go
RUN go build -o main .
 
# Etapa 2: Construcción de la imagen final
FROM python:3
 
RUN pip install awscli
 
WORKDIR /
 
# Copiar el binario compilado desde la etapa de compilación
COPY --from=builder /app/main .
 
# Exponer el puerto en el que corre la aplicación (cambiar al puerto adecuado)
EXPOSE 8080
 
# Comando para ejecutar la aplicación
CMD ["./main"]
 