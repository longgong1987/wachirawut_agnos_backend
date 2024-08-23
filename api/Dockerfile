FROM golang:1.22-alpine

WORKDIR /api

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3000

CMD ["air"]