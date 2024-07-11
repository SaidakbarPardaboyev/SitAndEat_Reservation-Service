package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	pb "reservation/genproto/resirvation"
	"reservation/pkg/logger"
	"reservation/storage/postgres"
)

type ReservationService struct {
	pb.UnimplementedResirvationServer
	Reser  *postgres.Reservation
	Logger *slog.Logger
}

func NewReservationService(db *sql.DB) *ReservationService {
	reser := postgres.NewReservationRepo(db)
	return &ReservationService{
		Reser:  reser,
		Logger: logger.NewLogger(),
	}
}

func (r *ReservationService) Createreservations(ctx context.Context, req *pb.RequestReservations) (*pb.Status, error) {
	resp, err := r.Reser.CreateReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni insert qilishda xatolik: %v", err))
		return nil, err
	}

	return resp, nil
}

func (r *ReservationService) GetAllReservations(ctx context.Context, req *pb.Void) (*pb.Reservations, error) {
	resp, err := r.Reser.GetAllReservation()
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetByIdReservations(ctx context.Context, req *pb.ReservationId) (*pb.Reservation, error) {
	resp, err := r.Reser.GetReservationByID(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Bitta malumotni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) UpdateReservations(ctx context.Context, req *pb.ReservationUpdate) (*pb.Status, error) {
	resp, err := r.Reser.UpdateReservations(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni update qilishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) DeleteReservations(ctx context.Context, req *pb.ReservationId) (*pb.Status, error) {
	resp, err := r.Reser.DeleteReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Resirvation deleting error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetReservationsByUserId(ctx context.Context, req *pb.UserId) (*pb.Reservations, error) {
	resp, err := r.Reser.GetReservationsByUserId(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("User id buyicha malumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) OrderMeal(ctx context.Context, req *pb.Order) (*pb.Status, error) {
	resp, err := r.Reser.OrderMeal(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Order service da xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) PayForReservation(ctx context.Context, req *pb.Payment) (*pb.Status, error) {
	resp, err := r.Reser.PayForReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Payment service da xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}
