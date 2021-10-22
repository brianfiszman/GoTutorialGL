FROM golang:1.17.1-alpine

# Add a work directory
WORKDIR /app

# Copy app files
COPY . .

RUN go mod download

# Run go app
CMD ["go","run","main.go"]
