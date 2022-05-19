FROM golang:1.18 as builder
WORKDIR /src
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/serve/main.go

FROM alpine as final
WORKDIR /opt/bin
COPY --from=builder /src/server .
ENTRYPOINT ["./server"]
