FROM golang:1.20-alpine3.19

RUN mkdir /alifE

WORKDIR /alifE

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go
CMD ["./main"]
