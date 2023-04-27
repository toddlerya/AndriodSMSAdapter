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

详细参数
```shell
启动服务

Usage:
  AndriodSMSAdapter server [flags]

Flags:
  -h, --help          help for server
      --mode string   运行模式:1. phone: 即为在手机内部运行 2. pc: 在PC上外部运行 (default "phone")
      --port uint16   设定服务端口 (default 30000)

Global Flags:
      --log-json           设置日志为JSON格式
  -l, --log-level string   设置日志级别 (default "INFO")
```


## 测试API服务

### 测试服务器启动状态

```shell
$ curl --location 'http://127.0.0.1:30000/'
{"message":"hi, I'm sms adapter"}
```

### 测试查询短信

```shell
curl --location 'http://127.0.0.1:30000/read' \
--header 'Content-Type: application/json' \
--data '{"where": "_id=220"} '

{
    "description": "查询短信信息",
    "message": "ok",
    "status_code": 200,
    "data": "Row: 0 _id=220, thread_id=32, address=10000, person=NULL, date=1602837752232, date_sent=1602837749000, protocol=0, read=1, status=-1, type=1, reply_path_present=0, subject=NULL, body=【公益短信】今天是世界粮食日，本周是我国粮食安全宣传周。让我们共同携手，厉行勤俭节约，反对餐饮浪费，把中国人的饭碗牢牢端在自己手上！国家粮食和物资储备局, service_center=+460030934772200, locked=0, error_code=0, seen=1, timed=0, deleted=0, sync_state=2, marker=128121671947195392, source=NTc1MDU0MjYwOTU3NjgyNTY6MDoxMDAwMDoxNjAyODM3NzUyMjMy, bind_id=57505426095768256, mx_status=0, mx_id=NULL, out_time=0, account=901235800, sim_id=1, block_type=0, advanced_seen=3, b2c_ttl=0, b2c_numbers=NULL, fake_cell_type=0, url_risky_type=0, creator=NULL, favorite_date=0, mx_id_v2=NULL, sub_id=-1\n"
}
```

### 更多复杂的查询语法

> 具体参见adb shell的语法

```shell
curl --location 'http://127.0.0.1:30000/read' \
--header 'Content-Type: application/json' \
--data '{"where": "body like '\''%同学您好，请于2022%'\''  and address='\''18952788xxx'\''", "sort": "date DESC"} '
```


## 原理

`--mode pc`时，执行`adb shell content query --uri content://sms/ --where "body like '%同学您好，请于2022%' " --sort "date DESC"`

`--mode phone`时【默认参数】, 执行`content query --uri content://sms/ --where "body like '%同学您好，请于2022%'" --sort "date DESC"`

注意

`--mode pc`不能使用sort、where的like条件and逻辑等复杂语法，`--mode phone`时可以