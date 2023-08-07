BUILD_DIR := build
EXAMPLES_DIR := $(wildcard server)
INCLUDE_PATH := $(abspath ../whisper.cpp)
LIBRARY_PATH := $(abspath ../whisper.cpp)

all: clean run

examples: $(EXAMPLES_DIR)


run: mkdir modtidy
	@echo Build $@
	@C_INCLUDE_PATH=${INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} go build ${BUILD_FLAGS} -o ${BUILD_DIR} ./server

mkdir:
	@echo Mkdir ${BUILD_DIR}
	@install -d ${BUILD_DIR}

modtidy:
	@go mod tidy

clean:
	@echo Clean
	@rm -fr $(BUILD_DIR)
	@go clean
