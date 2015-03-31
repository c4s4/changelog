NAME=changelog
VERSION=0.1.0
BUILD_DIR=build

build:
 mkdir -p $(BUILD_DIR)
 go build -o $(BUILD_DIR)/$(NAME)

run: build
	go run $(NAME).go

clean:
	rm -rf $(BUILD_DIR)
