SHELL:=/bin/bash
.PHONY: app_dist package build test clean

export GO111MODULE=on

BUILD_PATH=RELEASE_ANDROID_SMS_ADAPTER
TARGZ_NAME=$(BUILD_PATH).tar.gz
MD5_NAME=$(BUILD_PATH).md5

BUILD_ENV:=CGO_ENABLED=0
LDFLAGS:=-ldflags "-s -w "
BIN_PATH:=output/AndroidSMSAdapter
BUILD_SRC:=main.go
GOARCH=


clean:
	@ echo 清理打包发布目录和文件
	rm -rf output

before_build:
	mkdir -p output

build:
	@ echo 编译
	@ make before_build
	@ echo 开始编译
	$(BUILD_ENV) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(BIN_PATH) $(BUILD_SRC)
	@ make after_build

after_build: $(BIN_PATH)
	@ echo 二进制文件赋权
	@ ls -lh $(BIN_PATH)
	@ chmod +x $(BIN_PATH)


pre_package:
	@ echo 打包前准备
	mkdir -p $(BUILD_PATH)

app_dist:
	@ echo 准备应用打包文件
	cp $(BIN_PATH)                $(BUILD_PATH)
	tar czvf $(TARGZ_NAME) $(BUILD_PATH)
	md5sum  $(TARGZ_NAME) | awk '{ print $$1 }' > $(MD5_NAME)

package:
	make build
	make pre_package
	make app_dist

arm64_package:
	@ echo arm64架构打包
	make clean
	make package GOARCH=arm64