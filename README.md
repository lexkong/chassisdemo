# chassisdemo
go-chassis demo

# 微服务列表
* API： 服务对外接口
* RestServer： Rest服务
* GRpcServer： GRPC服务

# 生成协议文件
```
go generate protobuf/helloworld.go
```

# 修改生成的```helloworld.pb.go```
变量```_Greeter_serviceDesc```修改为```Greeter_serviceDesc```

# 编译
```go build```

# 运行
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  chassisdemo [command]

Available Commands:
  api         A brief description of your command
  help        Help about any command
  restserver  A brief description of your command
  grpcserver   A brief description of your command

Flags:
  -c, --config string   config file path (default is ./${command}/conf/)
  -h, --help            help for chassisdemo
  -t, --toggle          Help message for toggle

Use "chassisdemo [command] --help" for more information about a command.

```

### 运行API服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo api
```
> 测试时使用单实例

### 运行Rest服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo restserver
```
> 可以运行多个，运行多个时，可执行文件和config文件需要放到不同的目录，然后修改端口号

### 运行GRPC服务
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo grpcserver
```
> 可以运行多个，运行多个时，可执行文件和config文件需要放到不同的目录，然后修改端口号

# 测试
当所有服务都成功启动后，执行下面两个命令检查结果
```$xslt
curl localhost:5000/sayresthello/heavyrain
curl localhost:5000/saygrpchello/heavyrain
```

# 压力测试
当所有服务都成功启动后，执行下面两个命令进行性能验证
```$xslt
heavyrains-MacBook-Pro:chassisdemo heavyrainlee$ ./chassisdemo stress 0
```
> 0对grpc服务进行压力测试，1 对http服务压力测试。测试时使用100个协程，每个做1000次rpc，每次rpc休眠25ms