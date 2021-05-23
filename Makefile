export DB_CONNECTION=root:admin_petstore@tcp(127.0.0.1:3306)/petstore?parseTime=True&loc=UTC

start_service:
	go run src/main.go

create_db:
	docker run --name petStore_local -p '3306:3306' -e MYSQL_ROOT_PASSWORD=admin_petstore -e MYSQL_DATABASE=petstore -d mysql:8.0

.PHONY: start_service create_db