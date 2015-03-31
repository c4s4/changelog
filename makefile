NAME=changelog
VERSION=0.1.0
BUILD_DIR=build

run:
	go run $(NAME).go

clean:
	rm -rf $(BUILD_DIR)
