# Go 명령어 및 기본 설정
GO := go
APP_NAME := doppelganger-destroyer
CMD_DIR := cmd/$(APP_NAME)
BIN_DIR := bin
MOCK_DIR := mock/1

# 기본 빌드 경로 및 파일 이름
BUILD_OUTPUT := $(BIN_DIR)/$(APP_NAME)

# 기본 작업
.PHONY: all build run clean test lint

# 모든 작업 수행 (빌드 및 테스트)
all: build test

# 애플리케이션 빌드
build: clean
	@echo "==> 애플리케이션을 빌드하는 중..."
	mkdir -p $(BIN_DIR) # 빌드 결과물 디렉토리 생성
	$(GO) build -o $(BUILD_OUTPUT) $(CMD_DIR)/main.go

# 애플리케이션 실행
run: build
	@echo "==> 애플리케이션을 실행하는 중..."
	$(BUILD_OUTPUT) -dir $(MOCK_DIR)

# 테스트 실행
test:
	@echo "==> 테스트를 실행하는 중..."
	$(GO) test ./...

# 코드 린트 (golangci-lint 사용)
lint:
	@echo "==> 린트 검사를 실행하는 중..."
	golangci-lint run

# 빌드 결과물 제거
clean:
	@echo "==> 빌드 결과물을 정리하는 중..."
	rm -rf $(BIN_DIR)
