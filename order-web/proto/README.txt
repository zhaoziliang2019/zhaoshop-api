proto目录下执行如下语句
protoc -I . order.proto --go_out=plugins=grpc:.
protoc -I . goods.proto --go_out=plugins=grpc:.
protoc -I . inventory.proto --go_out=plugins=grpc:.