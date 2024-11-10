drestart:
	docker compose down && BUILDKIT_PROGRESS=plain docker compose build --no-cache && docker compose up -d

dlogs:
	docker compose logs -f

ddown:
	docker compose down

dup:
	docker compose up

dupb:
	BUILDKIT_PROGRESS=plain docker compose build --no-cache && docker compose up

dupd:
	docker compose up -d

dupdb:
	BUILDKIT_PROGRESS=plain docker compose build --no-cache && docker compose up -d

dps:
	docker compose ps