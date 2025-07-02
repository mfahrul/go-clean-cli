# go-clean-cli

⚠️ **EXPERIMENTAL PROJECT** ⚠️
This is an experimental project for generating CRUD boilerplate following clean architecture principles in Go. Please use with caution in production environments.

## Prerequisites

Before using this generator, ensure you have the following folder structure in your project:

```
your-project/
├── internal/
│   └── {domain}/           # e.g., user, product, etc.
│       ├── delivery/
│       │   └── http/      # HTTP handlers
│       ├── repository/    
│       │   └── postgres/  # Database implementations
│       ├── usecase/       # Business logic
│       └── test/          # Test files
├── pkg/                   # Shared packages
└── cmd/                   # Application entry points
```

## Features

### Code Generation
The generator creates the following files for your domain:

1. **Entity Layer**
   - `entity.go`: Domain models and business objects
   - `dto.go`: Data transfer objects for input/output

2. **Repository Layer**
   - `repository.go`: Repository interface definitions
   - `repositoryImpl.go`: Concrete implementation (PostgreSQL)

3. **UseCase Layer**
   - `usecase.go`: Business logic interface
   - `usecaseImpl.go`: Concrete implementation of business logic

4. **Delivery Layer**
   - `handler.go`: HTTP handlers with common CRUD operations
   - `route.go`: Route definitions and setup

5. **Testing**
   - `test.go`: Unit and integration test templates

### Built-in Features

#### 1. CRUD Operations
- Create new records
- Read single/multiple records with pagination
- Update existing records
- Delete records

#### 2. Pagination Support
The handler includes a robust pagination URL builder that:
- Uses Fiber's `c.Queries()` for query parameters
- Preserves existing query parameters
- Builds clean URLs for navigation
```go
// Example pagination URL building
queries := c.Queries()
var queryParts []string
for key, value := range queries {
    if key == "page" {
        continue
    }
    queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
}
```

#### 3. Clean Architecture Implementation
- Clear separation of concerns
- Dependency injection
- Interface-based design
- Easy to test and maintain

## Example Usage

1. **Set up your folder structure first:**
```bash
mkdir -p internal/user/delivery/http
mkdir -p internal/user/repository/postgres
mkdir -p internal/user/usecase
mkdir -p internal/user/test
```

2. **Run the generator:**
```bash
go run main.go generate user
```

3. **Generated files structure:**
```
internal/user/
├── delivery/
│   └── http/
│       ├── handler.go
│       └── routes.go
├── repository/
│   └── postgres/
│       └── user_repo.go
├── usecase/
│   └── user_usecase.go
├── test/
│   ├── integration_test.go
│   └── unit_test.go
├── dto.go
├── entity.go
├── repository.go
└── usecase.go
```

## Best Practices

1. **Entity Definition**
   - Define your entity struct in `entity.go`
   - Include validation tags
   - Add custom methods if needed

2. **Repository Implementation**
   - Implement custom queries in `repositoryImpl.go`
   - Use prepared statements for security
   - Handle database transactions properly

3. **UseCase Layer**
   - Add business logic in `usecaseImpl.go`
   - Validate input/output
   - Handle error cases

4. **HTTP Handlers**
   - Use proper HTTP status codes
   - Validate request payload
   - Implement proper error handling

## Contributing

This is an experimental project and contributions are welcome. Please ensure you:
1. Write tests for new features
2. Follow Go best practices
3. Document your changes
4. Keep the clean architecture principles in mind

## License

MIT License
