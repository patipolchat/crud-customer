# CRUD Customer

- Go web framework - Echo
- DB - SQLite
- DB ORM - GORM
- CMD - Cobra Cli

## Require
- Go1.22
- Docker - https://www.docker.com/
- Taskfile - https://taskfile.dev/
- Mockery - https://github.com/vektra/mockery

## Run Project
1. Create config.yaml
2. Migrate DB - `task auto-migrate`
3. Generate mockery - `task gen`
4. Serve Http Server with cobra command - `task serve-api`
5. Run test with coverage - `task test`

## Endpoint
- Create Customer - **POST - /api/v1/customers**
- Get All Customer - **GET - /api/v1/customers/**
- Get One Customer - **GET - /api/v1/customers/:id**
- Update Customer - **PUT - /api/v1/customers/:id**
- Delete Customer - **DELETE - /api/v1/customers/:id**


## config.yaml
```yml
server:
  port: 8080
  allowOrigins:
    - "*"
  bodyLimit: "10M" # MiB
  timeout: 30 # Seconds
  logLevel: DEBUG

database:
  file: "tmp/customer.db"
```
