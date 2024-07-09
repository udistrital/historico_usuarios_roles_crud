# Utilizar una imagen base de Go
FROM golang:1.18-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos de go.mod y go.sum y descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente de la aplicación
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que la aplicación escuchará
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./main"]