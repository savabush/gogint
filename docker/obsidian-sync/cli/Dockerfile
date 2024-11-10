FROM golang:1.22-alpine

WORKDIR /app

COPY obsidian-sync/go.mod obsidian-sync/go.sum ./

RUN go mod edit -replace github.com/savabush/lib=/app/lib

COPY obsidian-sync/cmd/obsidian-sync-cli/*.go ./
COPY obsidian-sync/internal ./internal
COPY lib/ ./lib

RUN go get github.com/savabush/lib && go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /obsidian-sync-cli

RUN rm -rf /app

CMD ["/obsidian-sync-cli"]