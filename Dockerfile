FROM golang:1.20 as build

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o myGolangApp

ENTRYPOINT ./myGolangApp