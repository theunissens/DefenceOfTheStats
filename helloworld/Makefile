
say_hello:
	@echo "Hello World"

build:
	@echo "Creating empty text files..."
	go tool cgo -exportheader gl/_cgo_export.h  gl/move.go
	go build

clean:
	@echo "Cleaning up..."
