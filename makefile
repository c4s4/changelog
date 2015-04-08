NAME=changelog
VERSION=0.1.0
BUILD_DIR=build
TEST_DIR=test

.PHONY: build test

test:
	go test

html: build
	$(BUILD_DIR)/$(NAME) to html $(TEST_DIR)/stylesheet.css < $(TEST_DIR)/CHANGELOG.yml > $(BUILD_DIR)/changelog.html

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(NAME)

install: build
	sudo cp $(BUILD_DIR)/$(NAME) /opt/bin/

clean:
	rm -rf $(BUILD_DIR)
