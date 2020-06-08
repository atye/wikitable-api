install:
	go mod tidy
	go install \
		github.com/golang/protobuf/protoc-gen-go \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \

protoc:
	protoc -I.  \
	-I$$(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway) \
	-I$$(go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway)/third_party/googleapis \
	--go_out=plugins=grpc,paths=source_relative:. \
	--grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true,paths=source_relative:. \
	--swagger_out=allow_merge=true,logtostderr=true:swagger/ \
	./service/pb/*.proto

test:
	go test -v -cover -race ./...

deploy:
	docker build -t wikitable-api .
	docker run --name wikitable-api-container -p 8080:8080 wikitable-api

down:
	docker stop wikitable-api-container
	docker rm wikitable-api-container