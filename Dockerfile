FROM golang:1.17.1-alpine
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod ./
RUN go mod download
# Copy app files
COPY . .
# Run go app
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -command="go run main.go"
