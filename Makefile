swag:
	swag init -g ./cmd/backend/main.go -o ./docs
check-db:
	docker compose exec postgres psql -U backenduser -d db
