package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	pb "reservation/genproto/restaurant"
	"reservation/pkg/logger"
	"reservation/storage/postgres"
)

type RestaurantService struct {
	pb.UnimplementedRestaurantServer
	Reser  *postgres.NewRestaurant
	Looger *slog.Logger
}

func NewRestaurantService(db *sql.DB) *RestaurantService {
	reser := postgres.NewRestaurantRepo(db)
	return &RestaurantService{
		Reser:  reser,
		Looger: logger.NewLogger(),
	}
}

func (s *RestaurantService) CreateRestaurant(ctx context.Context, req *pb.Restuarant) (*pb.Status, error) {
	resp, err := s.Reser.CreateRestaurant(req)
	if err != nil {
		s.Looger.Error(fmt.Sprintf("Error inserting data: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetAllRestaurants(ctx context.Context, req *pb.Void) (*pb.Restuanants, error) {
	resp, err := s.Reser.GetAllRestaurants(req)
	if err != nil {
		s.Looger.Error(fmt.Sprintf("Error getting all restaurants: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) GetRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.GetRes, error) {
	resp, err := s.Reser.GetRestaurant(req)
	if err != nil {
		s.Looger.Error(fmt.Sprintf("Error getting restaurant: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) UpdateRestaurant(ctx context.Context, req *pb.RestuarantUpdate) (*pb.Status, error) {
	resp, err := s.Reser.UpdateRestaurant(req)
	if err != nil {
		s.Looger.Error(fmt.Sprintf("Error updating restaurant: %v", err))
		return nil, err
	}
	return resp, nil
}

func (s *RestaurantService) DeleteRestaurant(ctx context.Context, req *pb.RestuanantId) (*pb.Status, error) {
	resp, err := s.Reser.DeleteRestaurant(req)
	if err != nil {
		s.Looger.Error(fmt.Sprintf("Error deleting restaurant: %v", err))
		return nil, err
	}
	return resp, nil
}
