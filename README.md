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
- [mockery](https://github.com/vektra/mockery) to generate mocks from interfaces

### Prerequisite
- local golang installation
- MySQL Database. Use Makefile to generate MySQL container locally.

## Development

- Starting Application

  The service expects database connection to be passed as an environment variable. In order to make local dev easy, there are two options from which one can run the service:

  - From Visual Studio Code:
    In order to debug the service, one can run the service directly from the Visual Studio Code `Debug and Run` > `Launch` and add breakpoints

  - Using Makefile:
    From the project Root:

    ```bash
    make start_service
    ```

- API End-points

  ```
  http://localhost:8080/docs/
  ```

- Running Unit Tests
  ```bash
  go test ./...
  ```