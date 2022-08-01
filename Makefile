build_server:
	docker compose build

start_server:
	docker compose up -d

stop_server:
	docker compose down

# Error codes: Success (0), Failure (1)
test_all:
	docker compose -f docker-compose.tests.yml up init-redis
	docker compose -f docker-compose.tests.yml up server_integration_test --build --exit-code-from server_integration_test

# Error codes: Success (0), Failure (1)
test_all_debug: 
	docker compose -f docker-compose.tests.yml up server_integration_test redis-insight init-redis -d --build