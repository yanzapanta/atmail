# Atmail Assessment Task
A Golang HTTP server that performs user management operations (CRUD) using RESTful APIs.

## Technologies & Packages Used:
- Go version 1.21.3
- Docker
- MySQL Docker Image: mysql:8.0
- Gin Web Framework
- Swagger API Documentation
- Wire for Dependency Injection
- Logrus for Logging

```
├── Dockerfile
├── cmd
│   ├── api
│       └── main.go
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── .env
├── go.mod
├── go.sum
├── makefile
├── internal
|   |── config
│       ├── db.go
│       |── env.go
│   ├── http
│   │   ├── route
│   │   ├── handler
│   │   ├── middleware
│   ├── wire
│   └── repository
│   └── service
│   └── mock
│   └── model
│   └── helper
└── resources
```

## Installation
1. Clone ```atmail``` repository
2. Run ```make up```
4. Localhost: ```http://localhost/atmail```
3. Swagger link:  ```http://localhost/atmail/swagger/docs/index.html```

## Endpoints
- [GET] /users - retrieves all users
- [POST] /users - creates a user
- [GET] /users/{id} - retrieves user details by ID
- [PUT] /users/{id} - Updates user details by ID
- [DELETE] /users/{id} - Deletes a user by ID

### Note: 
- Database ```atmail``` will be automatically created
- BasicAuth credentails:
    ```
        username: admin
        password: admin
    ```
- Refer to the ```makefile``` to see more commands