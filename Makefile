PKG_LIST := $(shell go list ./... | grep -v /packrd | grep -v /pkg/api-internal | grep -v /pkg/api-public | grep -v /mock | tr "\n" " ")
RELEASE_TAG = $(shell date -u +"%y%m%d_%H%M%S")

code-gen:
	rm -rf internal/server/protocol/http/api/v1/model
	mkdir -p internal/server/protocol/http/api/v1/model
	swagger generate model -f api/swagger/v1.yaml -t internal/server/protocol/http/api/v1 -m model

deps:
	go mod download

build:
	go build

run: migrate
	./Mytheresa service start

migrate: build
	./Mytheresa migrate up

validate_swagger:
	@IS_INVALID=`git status | grep api/swagger-ui/swagger.json | wc -l`; if [ $$IS_INVALID != 0 ]; then git diff api/swagger-ui/swagger.json; exit 1; fi
	@IS_INVALID=`git status | grep api/swagger-ui/internal-v1.json | wc -l`; if [ $$IS_INVALID != 0 ]; then git diff api/swagger-ui/internal-v1.json; exit 1; fi