FROM golang:alpine

RUN mkdir /app
WORKDIR /app
RUN apk add make

# Copy files from current directory to docker image
COPY . .

# Install wire for dependency injection
RUN go install github.com/google/wire/cmd/wire@latest
# Install Swagger for API documentation
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/githubnemo/CompileDaemon@latest

# Download project dependencies
RUN go mod tidy

# Generate swagger docs and run tests
RUN make swag && make test
# Build the application
RUN go build ./cmd/api/main.go

EXPOSE 80

ENTRYPOINT CompileDaemon --build="go build ./cmd/api/main.go" --command=./main