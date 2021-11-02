FROM golang:1.16

WORKDIR /gosample

COPY . .

# Install mockgen
RUN go install github.com/golang/mock/mockgen@v1.6.0

# Generate mock 
RUN go generate ./...

# Install dependencies
RUN go mod tidy
RUN go mod vendor

# Run test
RUN go clean -testcache && go test ./... -race

# Build service
RUN go build -o srv ./cmd/srv/...

RUN chmod +x ./srv
RUN chmod +x ./wait-for-it.sh

CMD ["./wait-for-it.sh" , "db:3306", "--strict" , "--timeout=30", "--", "sh","-c","./srv --http_port $PORT"]
