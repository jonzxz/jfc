.PHONY: default

define HELP
	make build
	make dependencies
	make run
	make addperson
	make addpayment
	make db (run db)
	make stopdb
endef

export HELP

default:
	@echo "$$HELP"

build:
	go build .

dependencies:
	go get .

run:
	go run .

addperson:
	curl localhost:8080/people/add --include --header "Content-Type: application/json" --request "POST" --data '{"name":"Jeff","TelegramId":"@jeff"}'

addpayment:
	curl localhost:8080/payments/add --include --header "Content-Type: application/json" --request "POST" --data '{"type":"Conservancy","remarks":"Conservancy for Dec", "totalamount":69.9}'

db:
	docker run --name="db" -d -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD='secretrootpassword' \
	-v ./sql:/docker-entrypoint-initdb.d \
	mysql/mysql-server:8.0.32

stopdb:
	docker container stop db && docker container rm db