include .env
AUTH_SERVER=auth
ORDER_SERVER=order

init:
	cd ${AUTH_SERVER} && go get .
	cd ${ORDER_SERVER} && go get .

build.auth:
	cd ${AUTH_SERVER} && docker build -t auth-server .

build.order:
	cd ${ORDER_SERVER} && docker build -t order-server .

build:
	make build.auth
	make build.order

update:
	-docker-compose down
	make build
	docker-compose up -d
