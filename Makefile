# Makefile: 一个工程中的源文件不计其数，其按类型、功能、模块分别放在若干个目录中，makefile定义了一系列的规则来指定哪些文件需要先编译，哪些文件需要后编译，哪些文件需要重新编译，甚至于进行更复杂的功能操作，因为 makefile就像一个Shell脚本一样，也可以执行操作系统的命令

# --- `protos`下级文件名的作为生成包名
# 无http网关
# package_name := helloctl
# 有http网关
package_name := helloctlgateway
# --- `protos`下级文件名的作为生成包名-end

# subst 函数用来文本替换，格式`$(subst from,to,text)`
out_package_name := $(strip $(subst -,, $(package_name)))

# wildcard 函数用来在 Makefile 中，替换 Bash 的通配符,格式`srcfiles := $(wildcard src/*.txt)`
proto_file := $(wildcard protos/$(package_name)/*.proto)

# 在指定目录生成grpc protocol buffers的go代码(生成序列化/反序列化和gRPC客户端/服务器代码)路径
go_out := ${PWD}/grpc_go/$(out_package_name)

# 在指定目录生成grpc protocol buffers的php代码(生成序列化/反序列化和gRPC客户端/服务器代码)路径
php_out := ${PWD}/grpc_php/$(out_package_name)

# protoc 命令参数说明
#	-I:指定要在其中搜索的目录;可以多次指定,且按顺序搜索目录;如果未指定则为当前工作目录
#	--go_out:在指定目录生成Go代码
#	–go_opt paths=source_relative:表示按源文件的目录组织输出,也就是" the same relative directory "。
#	关于XXX_opt 中 paths 有两个参数
#		- import(默认): 按照option go_package里面的那个路径生成
#		- source_relative：表示按源文件的目录组织输出。这意味着,会将proto文件相对proto_path(-I)指定的基目录,按同样的目录组织在`--go_out/--go-grpc_out/--grpc-gateway_out`上重放一遍
.PHONY: go_grpc go_grpc_gateway
go_grpc:
	@echo "[INFO] input  package name: $(package_name)"
	@echo "[INFO] output package name: $(out_package_name)"
	@echo "[INFO] go_out: $(go_out)"

	@rm -rf $(go_out)
	@mkdir -p $(go_out)

	@echo "[INFO] compiling"
	@protoc \
	-I . \
	--go_out $(go_out) \
	--go-grpc_out $(go_out) \
	--go-grpc_opt require_unimplemented_servers=false \
	$(proto_file)

	@echo "[INFO] compilation done"

# 指定生成grpc gateway 依赖的包`-I vendors`
go_grpc_gateway:
	@echo "[INFO] input  package name: $(package_name)"
	@echo "[INFO] output package name: $(out_package_name)"
	@echo "[INFO] go_out: $(go_out)"

	@rm -rf $(go_out)
	@mkdir -p $(go_out)

	@echo "[INFO] compiling"
	@protoc \
	-I . \
	-I vendors \
	--go_out $(go_out) \
	--go-grpc_out $(go_out) \
	--grpc-gateway_out $(go_out) \
	--go-grpc_opt require_unimplemented_servers=false \
	--grpc-gateway_opt logtostderr=true \
	$(proto_file)

	@echo "[INFO] compilation done"
