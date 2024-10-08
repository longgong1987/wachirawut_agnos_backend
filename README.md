# wachirawut_agnos_backend

## Required tools
- [VSCode](https://code.visualstudio.com)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [DBeaver](https://dbeaver.io/)

## Env (.env)
Need to create .env file in api directory

```bash
API_SECRET=
SERVER_PORT=3000
POSTGRESQL_HOST=postgres
POSTGRESQL_USERNAME=postgres
POSTGRESQL_PASSWORD=
POSTGRESQL_DATABASE=strong_password_step
POSTGRESQL_PORT=5432
```
## Deploy with docker compose
```bash
$ docker compose --env-file=api/.env up -d --build
```

## Stop and remove container
```bash
$ docker compose down
```

## Unit Test Coverage
```bash
$ cd api
$ go test ./internal/usecases/ -coverprofile=cover.out
$ go tool cover -html=cover.out
```
