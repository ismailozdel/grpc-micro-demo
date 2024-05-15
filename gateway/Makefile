.PHONY: modsync
modsync:
	go get -u
	go mod tidy
	go mod vendor

.PHONY: docker
docker: modsync
	docker build -t gateway-test .
