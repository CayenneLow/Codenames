build_server:
	docker compose build

start_server:
	docker compose up -d

stop_server:
	docker compose down

test_all:
	# Error codes: Success (0), Failure (1)
	docker compose -f docker-compose.tests.yml up init-redis
	docker compose -f docker-compose.tests.yml up server_integration_test --build --exit-code-from server_integration_test

test_all_debug: 
	# Error codes: Success (0), Failure (1)
	docker compose -f docker-compose.tests.yml up server_integration_test redis-insight init-redis -d --build