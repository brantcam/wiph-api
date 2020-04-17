run:
	go run server.go
docker-run:
	docker build -t wiph-api . && \
		docker run wiph-api