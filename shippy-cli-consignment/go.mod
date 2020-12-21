module github.com/siwonpawel/shippy/shippy-cli-consignment

go 1.15

require (
	github.com/siwonpawel/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.34.0
)

replace github.com/siwonpawel/shippy/shippy-service-consignment => ../shippy-service-consignment
