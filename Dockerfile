FROM golang:1.17.1-alpine

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod ./

# Copy envs
COPY .env ./

RUN go mod download
# Copy app files
COPY . .
# Run go app
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -command="env $(cat .env) go run main.go"
