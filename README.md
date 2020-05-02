# Wiremock [![Docker Pulls](https://img.shields.io/docker/pulls/prongbang/wiremock.svg)](https://hub.docker.com/r/prongbang/wiremock/)

> [Wiremock](https://hub.docker.com/r/prongbang/wiremock) Minimal Mock your APIs 

## Matching Routes using gorilla/mux

```
Read doc https://github.com/gorilla/mux#matching-routes
```

## Setup project

```shell script
project
├── docker-compose.yml
└── mock
    ├── login
    │   └── route.yml
    └── user
        ├── response
        │   └── user.json
        └── route.yml
```

#### Login

```shell script
POST http://localhost:8000/api/v1/login
```

- route.yml

```yaml
routes:
  login:
    request:
      method: "POST"
      url: "/api/v1/login"
    response:
      status: 200
      body: >
        {"message": "success"}
```

#### User

```shell script
GET   http://localhost:8000/api/v1/user/1
POST  http://localhost:8000/api/v1/user
```

- route.yml

```yaml
routes:
  get_user:
    request:
      method: "GET"
      url: "/api/v1/user/{id:[0-9]+}"
    response:
      status: 200
      body_file: user.json

  create_user:
    request:
      method: "POST"
      url: "/api/v1/user"
    response:
      status: 201
      body: >
        {"message": "success"}
```

## How to run

### Docker

```yaml
version: '3.7'
services:
  app_wiremock:
    image: prongbang/wiremock
    ports:
      - "8000:8000"
    volumes:
      - "./mock:/mock"
```

### Golang

```shell script
$ go get -u github.com/prongbang/wiremock
$ cd project
```

- Default port `8000`

```bash
$ wiremock
```

- Custom port `9000`

```bash
$ wiremock -port=9000
```

- Running

```shell script
  _      ___                        __  
 | | /| / (_)______ __ _  ___  ____/ /__
 | |/ |/ / / __/ -_)  ' \/ _ \/ __/  '_/
 |__/|__/_/_/  \__/_/_/_/\___/\__/_/\_\


 -> wiremock server started on :8000
```

### Example Project

[https://github.com/prongbang/wiremock-example](https://github.com/prongbang/wiremock-example)
