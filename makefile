run:
	go run caddy/main.go

test:
	k6 run --vus 1 --iterations 1 index.js

test-debug:
	DEBUG_ALL=true k6 run --vus 1 --iterations 1 index.js

tidy:
	go mod tidy

compile-proto:
	protoc --go_out=./caddy/entity/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=./caddy/entity/pb \
		--go-grpc_opt=paths=source_relative *.proto