NAME=changelog
VERSION=$(shell changelog release version)
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

archive: clean
	@echo "$(YELLOW)Building executable$(CLEAR)"
	mkdir -p $(BUILD_DIR)/$(NAME)-$(VERSION)/
	gox -output=$(BUILD_DIR)/$(NAME)-$(VERSION)/{{.Dir}}_{{.OS}}_{{.Arch}}
	cp LICENSE.txt $(BUILD_DIR)/$(NAME)-$(VERSION)/
	cp README.md $(BUILD_DIR)/ && cd $(BUILD_DIR) && md2pdf README.md && cp README.pdf $(NAME)-$(VERSION)/
	cd $(BUILD_DIR) && tar cvzf $(NAME)-bin-$(VERSION).tar.gz $(NAME)-$(VERSION)

release: test archive
	@echo "$(YELLOW)Releasing version $(VERSION)$(CLEAR)"
	release

clean:
	@echo "$(YELLOW)Cleaning generated files$(CLEAR)"
	rm -rf $(BUILD_DIR)

help:
	@echo "$(YELLOW)Print help$(CLEAR)"
	@echo "$(CYAN)test$(CLEAR)    Run unit tests"
	@echo "$(CYAN)build$(CLEAR)   Build executable"
	@echo "$(CYAN)archive$(CLEAR) Build binary archive"
	@echo "$(CYAN)release$(CLEAR) Make a release"
	@echo "$(CYAN)clean$(CLEAR)   Clean generated files"
	@echo "$(CYAN)help$(CLEAR)    Print this help screen"
