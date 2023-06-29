# ---- Builder----
FROM golang:1.20 as builder

WORKDIR /workspace
ENV GO111MODULE=on

COPY . .

RUN go mod tidy
RUN go build -o /bin/server


# ---- Stage ----
FROM debian:12-slim as stage

COPY --from=builder /bin/server /bin/server

ENTRYPOINT ["./bin/server"]