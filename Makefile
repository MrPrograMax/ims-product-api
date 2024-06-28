build:
	docker-compose build ims-product-api

run:
	docker-compose up ims-product-api

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up


