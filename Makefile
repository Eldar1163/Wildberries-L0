build:
		cd OrderServer && go build ./ && cd OrderPulisher && go build ./
run-server:
		cd OrderServer && go run ./
run-publisher:
		cd OrderPublisher && go run ./
test:
		cd OrderServer && go test ./...
docker-up:
		docker-compose up