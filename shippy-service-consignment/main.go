package main

import (
	"context"
	"log"

	micro "github.com/micro/go-micro/v2"
	pb "github.com/siwonpawel/shippy/shippy-service-consignment/proto/consignment"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {

	updated := append(repo.consignments, consignment)
	repo.consignments = updated

	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type consignmentService struct {
	repo repository
}

func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment

	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	consignments := s.repo.GetAll()
	res.Consignments = consignments

	return nil
}

func main() {

	repo := &Repository{}
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
	)

	service.Init()
	if err := pb.RegisterShippingServiceHandler(service.Server(), &consignmentService{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}

}
