## Install

> *Its works on windows, previously you should install make with choko*  
> *Link: https://stackoverflow.com/a/57042516*


#### Your Makefile declaration for generating protobuf
```shell
gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
  	    --go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative
```