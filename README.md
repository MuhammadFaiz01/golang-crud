# Person API

## Requirements

- Go version 1.22 or newer
- PostgreSQL version 15 or newer
- Dependency manager `go mod`
- Libraries:
  - [Gin](https://github.com/gin-gonic/gin) for web framework
  - [PGX](https://github.com/jackc/pgx) for PostgreSQL connection

## Project Structure

golang-crud/
├── config/
│   └── database.go         
├── controllers/
│   └── person_controller.go
├── models/
│   └── person.go           
├── routes/
│   └── person_routes.go    
├── .env                    
├── main.go                 
├── go.mod                  
