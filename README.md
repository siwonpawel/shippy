# shippy

Microservices in Golang using go-micro framework.

# How to run

Build images with commands executed from repository root folder

> docker build -f shippy-service-vessel\Dockerfile -t shippy-service-vessel .  
> docker build -f shippy-cli-consignment\Dockerfile -t shippy-cli-consignment .  
> docker build -f shippy-service-consignment\Dockerfile -t shippy-service-consignment .  

Run services

> docker run -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 shippy-service-consignment  
> docker run -d -e MICRO_SERVER_ADDRESS=:50051 shippy-cli-consignment  
> docker run -d -e MICRO_SERVER_ADDRESS=:50051 shippy-service-vessel  

For better outcome run **shippy-cli-consignment** multiple times and check output


# Credits

Created based on this [Microservices in Go serie](https://ewanvalentine.io/microservices-in-golang-part-0/)  
[Reference repository](https://github.com/EwanValentine/shippy)

# License

Under GPL-3.0 License