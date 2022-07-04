proto目录下执行如下语句
protoc -I . address.proto --go_out=plugins=grpc:.
protoc -I . message.proto --go_out=plugins=grpc:.
protoc -I . userfav.proto --go_out=plugins=grpc:.
protoc -I . goods.proto --go_out=plugins=grpc:.