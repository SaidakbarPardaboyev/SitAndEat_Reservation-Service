package service

import (
	"context"
	"database/sql"
	"log"

	pb "reservation/genproto/restaurant"
	"reservation/storage/postgres"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Reser *postgres.NewRestaurant
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	reser := postgres.NewRestaurantRepo(db)
	return &RestaurantService{
		Reser: reser,
	}
}

func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.Restuarant) (*pb.Status, error) {
	resp, err := s.Reser.CreateRestaurant(req)
	if err != nil {
		log.Println("Error inserting data:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetAllRestaurants(ctx context.Context, req *pb.Void) (*pb.Restuanants, error) {
	resp, err := s.Reser.GetAllRestaurants(req)
	if err != nil {
		log.Println("Error getting all restaurants:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.GetRes, error) {
	resp, err := s.Reser.GetRestaurant(req)
	if err != nil {
		log.Println("Error getting restaurant:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) UpdateRestaurant(ctx context.Context, req *pb.RestuarantUpdate) (*pb.Status, error) {
	resp, err := s.Reser.UpdateRestaurant(req)
	if err != nil {
		log.Println("Error updating restaurant:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) DeleteRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.Status, error) {
	resp, err := s.Reser.DeleteRestaurant(req)
	if err != nil {
		log.Println("Error deleting restaurant:", err)
		return nil, err
	}
	return resp, nil
}
