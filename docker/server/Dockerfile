FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/server/*.go ./
COPY internal ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-server

RUN rm -rf /app

EXPOSE 8080

CMD ["/youtubes-server"]