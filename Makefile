.PHONY: psql
psql:
	docker compose exec db psql "postgresql://jobs:supsupsup@localhost:5432/jobs"

.PHONY: migrate
migrate:
	docker compose exec app go run database/migrate.go
