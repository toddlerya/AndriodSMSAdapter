# AndriodSMSAdapter


## 环境准备

安卓手机打开开发者模式，允许调试，USB连接PC授权信任

## 运行调试

```shell
go mod tidy
go run main.go -l DEBUG --mode pc --port 30000
```

## 编译AMR二进制文件

```shell
make arm64_package
```

output目录下得到arm环境运行文件


## 手机内部运行

```shell
./AndroidSMSAdapter --port 30000
```

