<h1 align="center">
    Petstore-Openapi3
</h1>

<p align="center">
  <strong>Creating HTTP API using golang in openapi3 standard</strong><br>
</p>

## Contents
- [Getting started](#getting-started)
- [Development](#development)

## Getting started
The intent of this sample project is to demonstrate the usage of various tools to create webapis using openapi 3 standard. This project uses:
- [oapi-codegen](https://github.com/deepmap/oapi-codegen) to generate server stubs
- [swagger-ui](https://github.com/swagger-api/swagger-ui) to display the api document
- [echo](https://github.com/labstack/echo) as HTTP routing engine
- [statik](https://github.com/rakyll/statik) to compile static files into a Go binary

### Prerequisite
- local golang installation


## Development

- Starting Application
To start-up the application. From the project root, in the terminal:

```bash
go run src/main.go
```

- API End-points

```
http://localhost:8080/docs/
```