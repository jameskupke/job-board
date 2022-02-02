.PHONY: psql
psql:
	docker compose exec db psql "postgresql://jobs:supsupsup@localhost:5432/jobs"
