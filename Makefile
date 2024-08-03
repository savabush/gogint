build-server:
	go build -o youtubes-server cmd/cli/main.go

build-cli:
	go build -o youtubes-cli cmd/server/main.go

run-cli:
	go run cmd/cli/main.go

run-server:
	go run cmd/server/main.go

drestart:
	docker compose down && docker compose up -d --build

dlogs:
	docker compose logs -f

ddown:
	docker compose down

dup:
	docker compose up -d

dupd:
	docker compose up -d --build

dps:
	docker compose ps