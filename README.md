# API REST CON GO CRUD

#### NOTA
Todo se almacena en memoria

#### Puerto 
8080

#### Generar certificados
1. Abrir terminal en la raíz del proyecto
2. Ir a la capeta de certificates
3. Certificado privado (1024 hace referencia a la cantidad de bytes)
```
    openssl genrsa -out app.rsa 1024
```
4. Certificado publico 
```
    openssl rsa -in app.rsa -pubout > app.rsa.pub
```
## rutas
1. /v1/persons/create
2. /v1/persons/get-all
3. /v1/persons/update (requiere query param id)
4. /v1/persons/delete (requiere query param id)
5. /v1/persons/get-by-id (requiere query param id)
