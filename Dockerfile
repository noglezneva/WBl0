FROM golang:1.21-alpine AS builder

WORKDIR /build

ADD go.mod .

ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o main cmd/myapp/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

COPY /internal/view /build/view

CMD ["./main"]