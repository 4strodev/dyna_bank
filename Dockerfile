FROM golang:1.19

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o ./bin/app ./internal/cmd/server

ENTRYPOINT ["./bin/app"]
