module github.com/siwonpawel/shippy/shippy-service-consignment

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/siwonpawel/shippy/shippy-service-vessel v0.0.0-20201222212329-f927fc03eec5
	google.golang.org/grpc v1.29.1 // indirect
)
