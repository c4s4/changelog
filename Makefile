# Parent Makefiles https://github.com/c4s4/make

include ~/.make/Golang.mk

test: go-test # Run unit tests
release: go-release # Perform release
