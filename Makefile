V := @

OUT_DIR := ./target
OUT := $(OUT_DIR)/goapi
MAIN_PKG := ./cmd/goapi

default: build

.PHONY: build
build:
	@echo BUILDING $(OUT)
	$(V)go build -o $(OUT) $(MAIN_PKG)
	@echo DONE

.PHONY: vendor
vendor:
	$(V)go mod tidy -go=1.16 && go mod tidy -go=1.17
	$(V)go mod tidy
	$(V)go mod vendor

.PHONY: start
start:
	go run ./cmd/goapi/main.go

.PHONY: swagger
swagger:
	swagger generate spec -m -x vendor -o swagger-ui/swagger.json
