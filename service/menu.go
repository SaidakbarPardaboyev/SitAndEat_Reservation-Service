package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	pb "reservation/genproto/menu"
	"reservation/pkg/logger"
	"reservation/storage/postgres"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	Menu   *postgres.Menu
	Logger *slog.Logger
}

func NewMenuService(db *sql.DB) *MenuService {
	menuRepo := postgres.NewMenuRepo(db)
	return &MenuService{
		Menu:   menuRepo,
		Logger: logger.NewLogger(),
	}
}

func (m *MenuService) CreateFood(ctx context.Context, req *pb.CreateF) (*pb.Status, error) {
	resp, err := m.Menu.CreateFood(req)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Creating error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) GetAllFoods(ctx context.Context, req *pb.Void) (*pb.Foods, error) {
	resp, err := m.Menu.GetAllFoods()
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Malumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) GetFood(ctx context.Context, req *pb.FoodId) (*pb.Food, error) {
	resp, err := m.Menu.GetFood(req)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Malumotni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) UpdateFood(ctx context.Context, req *pb.UpdateF) (*pb.Status, error) {
	resp, err := m.Menu.UpdateFood(req)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Malumotni update qilishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) DeleteFood(ctx context.Context, req *pb.FoodId) (*pb.Status, error) {
	resp, err := m.Menu.DeleteFood(req)
	if err != nil {
		m.Logger.Error(fmt.Sprintf("Malumotni delete qilishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}
