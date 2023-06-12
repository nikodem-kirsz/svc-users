FROM golang:1.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

EXPOSE 8080

CMD ["server/grpc.go"]