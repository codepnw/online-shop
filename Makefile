postgresinit:
	docker run --name online-shop -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:15.4

postgres:
	docker exec -it online-shop psql

createdb:
	docker exec -it online-shop createdb online-shop

dropdb:
	docker exec -it online-shop dropdb online-shop