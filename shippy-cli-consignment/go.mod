module github.com/EwanValentine/shippy-cli-consignment

go 1.15

replace github.com/siwonpawel/shippy/shippy-service-consignment => ../shippy-service-consignment
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/micro/go-micro/v2 v2.9.1
	github.com/siwonpawel/shippy/shippy-service-consignment v0.0.0-20201221192730-00d93b578765
)