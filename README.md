# Senha Válida - Studio Sol

## How to Run

`URL: http://localhost:8080/verify`

### With Docker

```bash
docker build -t verifier-rest-service .
```

```bash
docker run --name verifier-rest -p 8080:8080 -d verifier-rest-service
```

### With Go

```bash
go run ./cmd/api
```

#### Tests
- Run all tests
```bash
go test -v ./cmd/api
```

- Run rules logic test
```bash
go test -v -run Test_rules ./cmd/api
```

- Run API test
```bash
go test -v -run Test_verifyPassword ./cmd/api
```