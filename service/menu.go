package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "reservation/genproto/menu"
	"reservation/storage/postgres"
	"reservation/pkg/logger"
)

type MenuService struct {
	pb.UnimplementedMenuServer
	db     *sql.DB
	menu   *postgres.Menu
	Logger *slog.Logger
}

func NewMenuService(db *sql.DB, menu *postgres.Menu) *MenuService {
	return &MenuService{
		db:     db,
		menu:   menu,
		Logger: logger.NewLogger(),
	}

}

func (m *MenuService) CreateFood(ctx context.Context, req *pb.CreateF) (*pb.Status, error) {
	resp, err := m.menu.CreateFood(req)
	if err != nil {
		m.Logger.Error("Creating error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) GetAllFoods(ctx context.Context, req *pb.Void) (*pb.Foods, error) {
	resp, err := m.menu.GetAllFoods(req)
	if err != nil {
		m.Logger.Error("Malumotlarni olishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) GetFood(ctx context.Context, req *pb.FoodId) (*pb.Food, error) {
	resp, err := m.menu.GetFood(req)
	if err != nil {
		m.Logger.Error("Malumotni olishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) UpdateFood(ctx context.Context, req *pb.UpdateF) (*pb.Status, error) {
	resp, err := m.menu.UpdateFood(req)
	if err != nil {
		m.Logger.Error("Malumotni update qilishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (m *MenuService) DeleteFood(ctx context.Context, req *pb.FoodId) (*pb.Status, error) {
	resp, err := m.menu.DeleteFood(req)
	if err != nil {
		m.Logger.Error("Malumotni delete qilishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}
