# ref: https://vic.demuzere.be/articles/golang-makefile-crosscompile/

BIN_DIR := build

default: clean frontend darwin linux windows integrity

clean:
	$(RM) $(BIN_DIR)/validate_cc*
	go clean -x

frontend:
	cd static && npm install && npm run build && cd ..

install:
	go install

darwin:
	GOOS=darwin GOARCH=amd64 go build -o '$(BIN_DIR)/validate_cc-darwin-amd64'

linux:
	GOOS=linux GOARCH=amd64 go build -o '$(BIN_DIR)/validate_cc-linux-amd64'

windows:
	GOOS=windows GOARCH=amd64 go build -o '$(BIN_DIR)/validate_cc-windows-amd64.exe'

integrity:
	cd $(BIN_DIR) && shasum *
