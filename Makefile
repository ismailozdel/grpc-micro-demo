gen:
	@protoc --go_out=. \
	--go_opt=module=github.com/ismailozdel/micro \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/ismailozdel/micro \
	proto/invoice/*.proto \
	proto/stock/*.proto \
	proto/user/*.proto


run-gateway:
	@cd gateway && go run .
	
run-invoice:
	@cd invoice && go run .

run-stock:
	@cd stock && go run .

run-user:
	@cd user && go run .

build-docker:
	@cd gateway && make docker && cd .. \
	&& cd invoice && make docker && cd .. \
	&& cd stock && make docker && cd .. \
	&& cd user && make docker && cd ..