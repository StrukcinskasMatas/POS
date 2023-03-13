ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
include $(ROOT_DIR)/.env

database:
	docker-compose up

run:
	go run main.go server
