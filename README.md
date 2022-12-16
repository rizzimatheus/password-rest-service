## How to Run

`URL: http://localhost:8080/verify`

### With Docker

```bash
docker build -t verifier-service .
```

```bash
docker run --name verifier-container -p 8080:8080 -d verifier-service
```

### With Go

```bash
go run ./cmd/api
```

