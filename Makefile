build:
	go build -o callback-server callback-server.go

run: build
	./callback-server

watch:
	reflex -s -r '\.go$$' make run 