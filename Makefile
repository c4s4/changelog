NAME=changelog
VERSION=0.1.0
BUILD_DIR=build
TEST_DIR=test

YELLOW=\033[1m\033[93m
CYAN=\033[1m\033[96m
CLEAR=\033[0m

.PHONY: build test

test:
	@echo "$(YELLOW)Running unit tests$(CLEAR)"
	go test

build:
	@echo "$(YELLOW)Building executable$(CLEAR)"
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(NAME)

release: test build
	@echo "$(YELLOW)Releasing version $(VERSION)$(CLEAR)"
	release ̀"`$(BUILD_DIR)/changelog release version`"

clean:
	@echo "$(YELLOW)Cleaning generated files$(CLEAR)"
	rm -rf $(BUILD_DIR)

help:
	@echo "$(YELLOW)Print help$(CLEAR)"
	@echo "$(CYAN)test$(CLEAR)    Run unit tests"
	@echo "$(CYAN)build$(CLEAR)   Build executable"
	@echo "$(CYAN)release$(CLEAR) Make a release"
	@echo "$(CYAN)clean$(CLEAR)   Clean generated files"
	@echo "$(CYAN)help$(CLEAR)    Print this help screen"
