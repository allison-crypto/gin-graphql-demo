STAGE?=local

ifneq (,$(wildcard ./.env))
	include .env
	export
endif

local-run:
	go run main.go
