# 生成 grpc 的 protocol buffer 协议的工具，及个人测试使用的公共包
// TODO: 1.安装命令。2.初始化go仓库。3.废弃https://github.com/leeprince/protobuf项目 - prince@todo 2022/11/28 上午11:22
---

# 一、gRPC
> 参考：https://grpc.io/docs/what-is-grpc/introduction/

## （一）gRPC
在gRPC中，客户端应用程序可以直接调用另一台计算机上的服务器应用程序上的方法，就好像它是本地对象一样，这使得您可以更轻松地创建分布式应用程序和服务。
与许多RPC系统一样，gRPC的基本思想是定义服务，指定可以通过其参数和返回类型远程调用的方法。
- 在服务器端，服务器实现此接口并运行gRPC服务器来处理客户端调用。
- 在客户端，客户端有一个存根（stub）（在某些语言中仅称为客户端），它提供与服务器相同的方法。

![gprc架构图](./doc/images/grpc.svg)

gRPC客户端和服务器可以在各种环境（从Google内部的服务器到您自己的桌面）中运行和相互通信，并且可以用gRPC支持的任何语言编写。
因此，例如，您可以轻松地用Java创建一个gRPC服务器，而客户端使用Go、Python或Ruby。
此外，最新的GoogleAPI将具有其接口的gRPC版本，使您可以轻松地将Google功能构建到应用程序中。

## （二）Protocol Buffers协议
默认情况下，gRPC使用Protocol Buffers协议（简称pb协议），这是Google成熟的开源机制，用于**序列化结构化数据**（尽管它可以用于其他数据格式，如JSON）。

使用协议缓冲区时，第一步是定义要在proto文件中序列化的数据的结构：这是一个扩展名为.proto普通文本文件。

协议缓冲区数据的结构为消息，其中每个消息都是一个小的信息逻辑记录，包含一系列名值对（称为字段）。

## （三）gRPC Go
> 参考: https://grpc.io/docs/languages/go/quickstart/

### 1.使用gRPC Go的前提条件
- Go最新三个主要版本中的任何一个。
- 协议缓冲(protocol buffer)编译器:[protoc](https://developers.google.com/protocol-buffers/docs/proto3)
- Go协议(protocol)编译器插件(安装gRPC Go protoc编译器插件)
    - 安装插件
    ```
        go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```
    - 添加到环境变量或者`$(go env GOBIN) `（确保$GOBIN位于$PATH中。）
    ```
        export PATH="$PATH:$(go env GOBIN)"
    ```

### 2.项目中使用
> 参考: https://github.com/grpc/grpc-go
- 引入包
```
import "google.golang.org/grpc"
```

- FAQ

(1)`go get google.golang.org/grpc` 可能下载失败;用下面这个方式代替

报错
```
$ go get -u google.golang.org/grpc
package google.golang.org/grpc: unrecognized import path "google.golang.org/grpc" (https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
```
解决
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc 
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto cd $GOPATH/src/

go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest
go mod tidy
go mod vendor
go build -mod=vendor
```

## （四）gRPC Go gateway
> 参考：https://github.com/grpc-ecosystem/grpc-gateway
### 1. 安装gRPC Go gateway protoc编译器插件
(1)安装可执行文件，如果未下载包则会自动下载
```
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
(2)这将在您的$GOBIN中放置四个二进制文件(确保$GOBIN在$PATH中)
```
    protoc-gen-grpc-gateway
    protoc-gen-openapiv2
    protoc-gen-go
    protoc-gen-go-grpc
```

## （五）micro v3 相关包
> https://github.com/go-micro/go-micro/
```
go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3
```

# 二、生成协议缓冲区（Protocol Buffers）协议的protoc工具

## （一）使用协议缓冲区语言构建协议缓冲区数据
[.proto文件语法以及如何从.proto文件生成数据访问类](https://developers.google.cn/protocol-buffers/docs/proto3)

## （二）安装protoc
- 到 https://github.com/protocolbuffers/protobuf/releases/ 下载相应系统的 protoc 文件
> 本人系统是：macos, 下载的是：protoc-3.17.3-osx-x86_64.zip

- 解压后，将 protoc（protoc-3.17.3-osx-x86_64/bin/protoc） 复制到 $GOPATH/bin 路径下

## （三）安装protoc相关命令
### 1. gRPC Go protoc编译器插件
请看`一、gRPC` -> `gRPC Go` -> `1.使用gRPC Go的前提条件` -> `- Go协议(protocol)编译器插件(protoc命令的Go协议编译器插件)`
### 1. gRPC Go gateway protoc编译器插件
请看`一、gRPC` -> `（四）gRPC Go gateway` -> `1. 安装gRPC Go gateway protoc编译器插件`

## （四）protobuf协议（Protocol Buffers）应存放的位置
1. 【推荐】放到公共仓库（github、gitlab）中
2. 放到$(GOPATH)
3. 放到当前项目中

# 三、Makefile
一个工程中的源文件不计其数，其按类型、功能、模块分别放在若干个目录中，makefile定义了一系列的规则来指定哪些文件需要先编译，哪些文件需要后编译，哪些文件需要重新编译，甚至于进行更复杂的功能操作，因为 makefile就像一个Shell脚本一样，也可以执行操作系统的命令

# 四、目录结构说明
```
.
├── Makefile # Makefile 文件
├── README.md # 项目必读手册
├── pbapidoc # Makefile 文件中`pb_api_doc`命令生成的gRPC protobuff（protocol buffers）协议的文档
├── bin # 所有的可执行文件。注：大部分命令安装在操作系统中，未直接使用该路径下的可执行文件
├── doc # 文档相关资源
├── go.mod # golang的包管理
├── go.sum # golang的包校验
├── grpc_go # Makefile 文件中`go_grpc`命令生成的golang gRPC协议文档
│   ├── helloctl # 生成的指定包名
│   └── helloctlgateway # 生成的指定包名
│       ├── hello.pb.go # 通过`protoc-gen-go`插件生成
│       ├── service_helloctlgateway.pb.go # 通过`protoc-gen-go`插件生成
│       ├── service_helloctlgateway.pb.gw.go # 通过`protoc-gen-grpc-gateway`插件生成
│       └── service_helloctlgateway_grpc.pb.go # 通过`protoc-gen-go-grpc`插件生成
├── grpc_php # Makefile 文件中`php_grpc`命令生成的php gRPC协议文档
│   └── helloctl # 生成的指定包名
│       ├── GPBMetadata # 生成一个对应包名的.php类文件，用于保存.proto的二进制元数据
|       └── Helloctl
|           ├── HelloCtlClient.php # php gRPC客户端类
|           ├── HelloCtlStub.php # php gRPC服务端类
|           ├── SayHelloReq.php # php gRPC的请求体类
|           ├── SayHelloResp.php # php gRPC的响应体类
|           └── SayHelloRespData.php # php gRPC的响应体的数据类
├── protos # .proto 文件定义目录
│   ├── helloctl # 包/命名空间
│   │   ├── _helloctl.proto # 定义 rpc 的.proto文件
│   │   └── hello.proto # 定义请求、响应的消息体结构的.proto文件
│   └── helloctlgateway # 包/命名空间。包含网关层，如golang的gPRC gateway
│       ├── _helloctl.proto # 定义 rpc 的.proto文件
│       └── hello.proto # 定义请求、响应的消息体结构的.proto文件
└── vendors # protoc 命令依赖的包
```

# 五、使用
## （一）定义`.proto`文件
### 1. cd 到Makefile 所在目录
cd 到Makefile 所在目录中，当前Makefile所在目录为项目根目录

### 2. 创建`项目名称`
在 `protos` 目录下，根据项目创建`项目名称`

> 项目名称的规则（参考：golang包名的定义）：去除所有非字母或者数字的字符，再转成全小写）目录

### 3. 开始编写`.proto`
> `service_{项目名称}.proto`为定义rpc服务的文件

（1）定义`package`, 与`项目名称`一样
（2）根据golang或php语言定义相关option
    > golang: `// option go_package = "{path};{go package名称}"; - {path}:会自动生成`
    > php: 暂时没用到
（3）按照API设计，结合[proto语法](https://developers.google.cn/protocol-buffers/docs/proto3)编写`.proto`文件

## （二）使用Makefile中定义的命令
### 1. 生成golang 的gRPC protobuff（protocol buffers）协议
```
# 当前Makefile中默认的命令即为：go_grpc，所以直接执行：`make` 效果一样
$ make go_grpc
```

### 2. 生成php 的gRPC protobuff（protocol buffers）协议
```
$ make php_grpc
```

### 3. 生成`.proto`中定义的gRPC protobuff（protocol buffers）协议的文档
```
$ make pb_api_doc
```

### 4. 安装生成gRPC protobuff（protocol buffers）协议所需的可执行文件工具
```
$ make tools
```