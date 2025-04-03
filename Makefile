docker:
	docker compose --env-file .env -f deployment/docker-compose.yaml up -d --build
