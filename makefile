NAME=changelog
VERSION=0.1.0
BUILD_DIR=build
TEST_DIR=test

.PHONY: build test

test:
	go test

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(NAME)

release: test build
	release ̀"`$(BUILD_DIR)/changelog release version`" "`$(BUILD_DIR)/changelog release summary`"

clean:
	rm -rf $(BUILD_DIR)
