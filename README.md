

# Simple Key-Value Store With Go

In memory key value store application using go

![Website](https://img.shields.io/website?url=http%3A%2F%2F206.81.16.177%3A8080%2Fget%3Fkey) 

![Docker Image Size (tag)](https://img.shields.io/docker/image-size/yalcinc/simple-kvs-go/latest)

![Docker Image Version (tag latest semver)](https://img.shields.io/docker/v/yalcinc/simple-kvs-go/latest)
## Get Started

`cp config .env` 

`go build main.go`
## Architecture

```
in-memory-key-value-store
├── domain
│   ├── record.go
│   ├── store_test.go
│   └── store.go
├── infrastructure
│   ├── env.go
│   ├── logger.go
│   ├── router.go
│   ├── handler_test.go
│   └── handler.go
├── interfaces
│   ├── handler.go
│   ├── store_controller.go
│   └── store_repository.go
├── log
│   ├── server.log
│   └── error.log
├── temp
│   └── kvs_data.json
├── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── usecases
    ├── logger.go
    ├── store_interactor.go
    └── store_repository.go

```
  
## API

#### Add Record

```http
  POST /add?key={key}&value={value}
```

| Parameter | Type     | Explanation                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Required**.
 | `value`      | `string` | **Required**. 



#### Get Value

```http
  GET /get?key={key}
```

| Parameter | Type     | Explanation                |
| :-------- | :------- | :------------------------- |
| `key` | `string` | **Required** Key of the value |

#### Set Value

```http
  PUT /set?key={key}
```

| Parameter | Type     | Explanation                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `string` | **Required**. Key of the value to be changed
 


  
## Controller

| Controller Method | HTTP Method | Description                                |
|-------------------|-------------|--------------------------------------------|
| AddRecord             | POST         | Adds a new record.         |
| GetValue             | GET        | Returns the value of the passed key |
| SetValue              | PUT         | Changes the value of the passed key     |

## Live Demo

[Demo](http://206.81.16.177:8080/) 


  
## Environment Variables

To run this project you will need to add the following environment variables to your .env file

`BACKUP_TIME`

`PORT`

  
