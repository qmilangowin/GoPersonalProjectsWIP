.PHONY: all build_wasm build_admin

all: build_wasm build_admin

build_wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm

build_admin:
	go build -o ./admin -tags=embed .

.PHONY:
run_app:
	./admin