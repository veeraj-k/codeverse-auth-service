FROM golang:1.24-alpine AS base


WORKDIR /build


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o codeverse-auth-svc


EXPOSE 4000

CMD ["./codeverse-auth-svc"]
