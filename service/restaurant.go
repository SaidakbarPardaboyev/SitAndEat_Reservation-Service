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
	db    *sql.DB
	reser *postgres.NewRestaurant
}

func NewRestaurantService(db *sql.DB, reser *postgres.NewRestaurant) *RestaurantService {
	return &RestaurantService{db: db, reser: reser}
}

func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.Restuarant) (*pb.Status, error) {
	resp, err := s.reser.CreateRestaurant(req)
	if err != nil {
		log.Println("Error inserting data:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetAllRestaurants(ctx context.Context, req *pb.Void) (*pb.Restuanants, error) {
	resp, err := s.reser.GetAllRestaurants(req)
	if err != nil {
		log.Println("Error getting all restaurants:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.GetRes, error) {
	resp, err := s.reser.GetRestaurant(req)
	if err != nil {
		log.Println("Error getting restaurant:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) UpdateRestaurant(ctx context.Context, req *pb.RestuarantUpdate) (*pb.Status, error) {
	resp, err := s.reser.UpdateRestaurant(req)
	if err != nil {
		log.Println("Error updating restaurant:", err)
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) DeleteRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.Status, error) {
	resp, err := s.reser.DeleteRestaurant(req)
	if err != nil {
		log.Println("Error deleting restaurant:", err)
		return nil, err
	}
	return resp, nil
}
