FROM golang:alpine as builder

RUN mkdir -p /print-solution/server
WORKDIR /print-solution/server
COPY go.mod .
COPY go.sum .

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine:latest

RUN mkdir -p /cmd/static
WORKDIR /cmd

COPY --from=builder /print-solution/server .

CMD ["./server"]