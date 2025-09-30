FROM golang:1.24.6-alpine3.22 as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -a -installsuffix cgo -o main cmd/main.go
WORKDIR /dist
RUN cp /build/main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /dist/main .
EXPOSE 8080
CMD ["./main"]