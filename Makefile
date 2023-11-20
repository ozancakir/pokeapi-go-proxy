build:
	go build -ldflags "-w" -o ./bin/$(APP_NAME) ./cmd/*.go

run:
	go run  ./cmd/*.go

