# Api-Libros Simple

Sin base de datos.

## Para Empezar.    
``` bash
# Instalar mux Router
go get -u github.com/gorilla/mux
```

``` bash
go build
./api-books
```

``` bash

http://localhost:8000/books

```

## Endpoints
Atravez del Postman

### Obtener todos los libros
``` bash
GET api/books
```

### Obtener un solo libro
``` bash
GET api/books/{id}
```

### Borrar un libro
``` bash
DELETE api/books/{id}
```

### Crear un Libro
``` bash
POST api/books

# Request sample
# {
#   "isbn":"123456",
#   "title":"Libro 3",
#   "author":{"firstname":"Robert",  "lastname":"Jone"}
# }
```

### Actualizar un Libro
``` bash
PUT api/books/{id}

# Request sample
# {
#   "isbn":"123456",
#   "title":"Libro 3",
#   "author":{"firstname":"Robert",  "lastname":"Jone"}
# }

```

