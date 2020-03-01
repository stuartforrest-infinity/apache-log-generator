FROM golang:1.13 as builder
WORKDIR /go/src/github.com/fergusstrange/apache-log-generator
COPY main.go .
COPY go.mod .
COPY go.sum .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:3.9.5
WORKDIR /root
COPY --from=builder /go/src/github.com/fergusstrange/apache-log-generator/main .
ENTRYPOINT ["./main"]