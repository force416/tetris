APP_NAME := Tetris
MODULE := ./cmd/tetris
DIST_DIR := dist
BUILD_DIR := $(DIST_DIR)/build
APP_DIR := $(DIST_DIR)/$(APP_NAME).app
APP_CONTENTS := $(APP_DIR)/Contents
APP_MACOS := $(APP_CONTENTS)/MacOS
APP_RESOURCES := $(APP_CONTENTS)/Resources
TARGET_OS := darwin
TARGET_ARCH ?= $(shell go env GOARCH)
VERSION ?= 0.1.0
BIN_PATH := $(BUILD_DIR)/$(APP_NAME)
PLIST_PATH := $(APP_CONTENTS)/Info.plist
DMG_PATH := $(DIST_DIR)/$(APP_NAME)-$(TARGET_OS)-$(TARGET_ARCH).dmg

.PHONY: help run test build clean app bundle plist dmg

help:
	@echo "可用目標："
	@echo "  make run    - 啟動視窗版 Tetris"
	@echo "  make test   - 執行 go test ./..."
	@echo "  make build  - 編譯執行檔到 $(BIN_PATH)"
	@echo "  make app    - 打包 macOS .app 到 $(APP_DIR)"
	@echo "  make dmg    - 打包 macOS .dmg 到 $(DMG_PATH)"
	@echo "  make clean  - 清除 dist 目錄"

run:
	go run $(MODULE)

test:
	go test ./...

build: $(BIN_PATH)

$(BIN_PATH):
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=1 GOOS=$(TARGET_OS) GOARCH=$(TARGET_ARCH) go build -o $(BIN_PATH) $(MODULE)

app: clean $(APP_MACOS)/$(APP_NAME) $(PLIST_PATH)

bundle: app

$(APP_MACOS)/$(APP_NAME): $(BIN_PATH)
	mkdir -p $(APP_MACOS)
	mkdir -p $(APP_RESOURCES)
	cp $(BIN_PATH) $(APP_MACOS)/$(APP_NAME)
	chmod +x $(APP_MACOS)/$(APP_NAME)

$(PLIST_PATH):
	mkdir -p $(APP_CONTENTS)
	printf '%s\n' \
	'<?xml version="1.0" encoding="UTF-8"?>' \
	'<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">' \
	'<plist version="1.0">' \
	'<dict>' \
	'    <key>CFBundleDevelopmentRegion</key>' \
	'    <string>zh_TW</string>' \
	'    <key>CFBundleExecutable</key>' \
	'    <string>$(APP_NAME)</string>' \
	'    <key>CFBundleIdentifier</key>' \
	'    <string>com.ericshih.tetris</string>' \
	'    <key>CFBundleInfoDictionaryVersion</key>' \
	'    <string>6.0</string>' \
	'    <key>CFBundleName</key>' \
	'    <string>$(APP_NAME)</string>' \
	'    <key>CFBundlePackageType</key>' \
	'    <string>APPL</string>' \
	'    <key>CFBundleShortVersionString</key>' \
	'    <string>$(VERSION)</string>' \
	'    <key>CFBundleVersion</key>' \
	'    <string>1</string>' \
	'    <key>LSMinimumSystemVersion</key>' \
	'    <string>13.0</string>' \
	'    <key>NSHighResolutionCapable</key>' \
	'    <true/>' \
	'</dict>' \
	'</plist>' > $(PLIST_PATH)

dmg: app
	hdiutil create -volname "$(APP_NAME)" -srcfolder "$(APP_DIR)" -ov -format UDZO "$(DMG_PATH)"

clean:
	rm -rf $(DIST_DIR)
