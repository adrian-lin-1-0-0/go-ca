codegen:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
    -i /local/api/openapi-spec/v2.yaml \
    -l go-server \
    -o /local/swagger-codegen
run:
	go run cmd/app/main.go
