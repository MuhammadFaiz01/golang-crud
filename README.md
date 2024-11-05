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
1. Endpoint `POST`, `/person`
   ```json
    {
    "full_name": "John Doe",
    "age": 30,
    "birth_date": "1993-04-15",
    "address": "123 Main St, Cityville"
    }
    ```
- Status `201 Created`
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
