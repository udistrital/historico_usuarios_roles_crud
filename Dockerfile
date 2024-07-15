# Utilizar una imagen base de Go
#FROM golang:1.18-alpine

# Establecer el directorio de trabajo dentro del contenedor
#WORKDIR /app

# Copiar los archivos de go.mod y go.sum y descargar las dependencias
#COPY go.mod go.sum ./
#RUN go mod download

# Copiar el resto del código fuente de la aplicación
#COPY ./. .

# Compilar la aplicación
#RUN go build -o main .

# Exponer el puerto en el que la aplicación escuchará
#EXPOSE 8080

# Ejecutar la aplicación
#CMD ["./main"]
##############################
# Usa una imagen base de Go
FROM golang:1.18

# Añade GOPATH/bin al PATH
ENV PATH="/go/bin:${PATH}"

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum desde el directorio raíz del proyecto
COPY go.mod go.sum

# Descarga las dependencias de Go
#RUN go mod download

# Copia todos los archivos del proyecto desde el directorio raíz del proyecto
COPY . .

# Instala bee utilizando go install
RUN go install github.com/beego/bee@v1.12.3
#RUN bee version
# Install any needed dependencies specified in go.mod
#RUN go mod tidy

# Expone el puerto en el que la aplicación Beego estará escuchando
EXPOSE ${API_PORT}

# Comando para ejecutar la aplicación
CMD ["./main"]
