NAME=changelog
VERSION=0.1.0
BUILD_DIR=build

.PHONY: build test

test:
	go test

html: build
	$(BUILD_DIR)/$(NAME) to html test/stylesheet.css > $(BUILD_DIR)/changelog.html

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(NAME)

run: build
	$(BUILD_DIR)/$(NAME)

clean:
	rm -rf $(BUILD_DIR)
