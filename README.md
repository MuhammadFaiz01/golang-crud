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
      "full_name": "Faiz",
      "age": 31,
      "birth_date": "1992-04-15",
      "address": "456 New Street, City"
    }
    ```
- Status `201 Created`
- Body
  ```json
    {
      "id": 0,
      "full_name": "Faiz",
      "age": 31,
      "birth_date": "1992-04-15",
      "address": "456 New Street, City"
    }
  ```

  **Get a Person by Full Name**
  - Endpoint `GET`, `/person?full_name=Sabit`
  - Contoh Request `http GET /person?full_name=Faiz`
  - Status `200 ok`
  - Body
     ```json
    {
        "id": 1,
        "full_name": "John Doe",
        "age": 30,
        "birth_date": "1993-04-15",
        "address": "123 Main St, Cityville"
    }
  ```
**Update an Existing Person**
- Endpoint `PUT`, `PUT /person?id=1`
   ```json
    {
      "full_name": "Aflah",
      "age": 20,
      "birth_date": "2005-10-15",
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
    ```json
    {
      DELETE /person?id=1
    }
    ```
- Status `200 ok`
- - Body
     ```json
    {
        "message": "Person deleted successfully"
    }
  ```
