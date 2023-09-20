# Utilizamos una imagen base de Go
FROM golang:latest

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos el código fuente al directorio de trabajo en el contenedor
COPY . .

# Compilamos la aplicación Go
RUN go build -o main .

# Exponemos el puerto 8080 en el contenedor
EXPOSE 8080

# Comando para ejecutar la aplicación cuando se inicie el contenedor
CMD ["./main"]
