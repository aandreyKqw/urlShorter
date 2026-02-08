include .env
export 

service-run:
	@cd cmd && go run .

migrate-up:
	@cd iternal/ && \
	migrate -path migrations -database ${DB_CONN} up

migrate-down:
	@cd iternal/ && \
	migrate -path migrations -database ${DB_CONN} down

migrate-version:
	@cd iternal/ && \
	migrate -path migrations -database ${DB_CONN} version