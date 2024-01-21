V := @

NAME := goapi
OUT_DIR := ./target

MAIN_OUT := $(OUT_DIR)/$(NAME)
MAIN_PKG := ./cmd/$(NAME)goapi

default: build

.PHONY: build
build:
	@echo BUILDING $(MAIN_OUT)
	$(V)go build -o $(MAIN_OUT) $(MAIN_PKG)
	@echo DONE

.PHONY: run
run:
	go run ./cmd/$(NAME)/main.go

.PHONY: vendor
vendor:
	$(V)go mod tidy -go=1.16 && go mod tidy -go=1.17
	$(V)go mod tidy
	$(V)go mod vendor

.PHONY: swagger
swagger:
	swagger generate spec -m -x vendor -o swagger-ui/swagger.json
