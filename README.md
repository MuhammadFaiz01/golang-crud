# Person API

## Requirements

- Go version 1.22 or newer
- PostgreSQL version 15 or newer
- Dependency manager `go mod`
- Libraries:
  - [Gin](https://github.com/gin-gonic/gin) for web framework
  - [PGX](https://github.com/jackc/pgx) for PostgreSQL connection

## Project Structure

.
├── ...
├── test                    # Test files (alternatively `spec` or `tests`)
│   ├── benchmarks          # Load and stress tests
│   ├── integration         # End-to-end, integration tests (alternatively `e2e`)
│   └── unit                # Unit tests
└── ...
