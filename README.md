# Person API

## Requirements

- Go VERSI 1.22 atau yang paling baru
- PostgreSQL versi 15 atau yang paling baru
- Dependency manager `go mod`
- Libraries:
  - [Gin](https://github.com/gin-gonic/gin)
  - [PGX](https://github.com/jackc/pgx)
  - [GODOTENV](https://github.com/joho/godotenv)

## Setup Instructions
**Clone the Repository**
1.  ```bash git clone https://github.com/MuhammadFaiz01/golang-crud.git```
2. cd person-api
3. install dependency
4. go run main.go

## API Endpoints and Testing
**Create New Person**
- Endpoint `POST`, `/person`
   ```json
    {
      "full_name": "rendi",
      "age": 25,
      "birth_date": "1994-05-15T00:00:00Z",
      "address": "Jalan Kemang"
    }
    ```
- Status `201 Created`
- Body
  ```json
    {
      "id": 0,
      "full_name": "Faiz",
      "age": 31,
      "birth_date": "1994-05-15T00:00:00Z",
      "address": "456 New Street, City"
    }
  ```

**Get a Person by Full Name**
  - Endpoint `GET`, `/person?full_name=full_name`
  - Contoh Request `http GET /person?full_name=Indra`
  - Status `200 ok`
  - Body
     ```json
    {
        "id": 5,
        "full_name": "Indra",
        "age": 19,
        "birth_date": "2005-10-15T00:00:00Z",
        "address": "Jalan Rambang"
    }
    ```
**Update an Existing Person**
- Endpoint `PUT`, `PUT /person?id=1`
   ```json
    {
      "full_name": "Aflah",
      "age": 20,
      "birth_date": "1994-05-15T00:00:00Z",
      "address": "Jalan Rambang"
    }
    ```
- Status `200 ok`
- Body
     ```json
    {
     "message": "Person updated successfully"
    }
  ```
     
**Delete a Person by ID**
- Endpoint `DELETE`, `DELETE /person?id=${id}`
    ```URL
      DELETE /person?id=1
    ```
- Status `200 ok`
- - Body
     ```json
    {
        "message": "Person deleted successfully"
    }
    ```
