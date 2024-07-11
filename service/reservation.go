package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "reservation/genproto/resirvation"
	"reservation/pkg/logger"
	"reservation/storage/postgres"
)

type ReservationService struct {
	pb.UnimplementedResirvationServer
	db     *sql.DB
	reser  *postgres.Reservation
	Logger *slog.Logger
}

func NewReservationService(db *sql.DB, reser *postgres.Reservation) *ReservationService {
	return &ReservationService{
		db:     db,
		reser:  reser,
		Logger: logger.NewLogger(),
	}
}

func (r *ReservationService) Createreservations(ctx context.Context, req *pb.RequestReservations) (*pb.Status, error) {
	resp, err := r.reser.CreateReservation(req)
	if err != nil {
		r.Logger.Error("Malumotlarni insert qilishda xatolik: %v", err)
		return nil, err
	}

	return resp, nil
}

func (r *ReservationService) GetAllReservations(ctx context.Context, req *pb.Void) (*pb.Reservations, error) {
	resp, err := r.reser.GetAllReservation()
	if err != nil {
		r.Logger.Error("Malumotlarni olishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetByIdReservations(ctx context.Context, req *pb.ReservationId) (*pb.Reservation, error) {
	resp, err := r.reser.GetReservationByID(req)
	if err != nil {
		r.Logger.Error("Bitta malumotni olishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) UpdateReservations(ctx context.Context, req *pb.ReservationUpdate) (*pb.Status, error) {
	resp, err := r.reser.UpdateReservations(req)
	if err != nil {
		r.Logger.Error("Malumotlarni update qilishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) DeleteReservations(ctx context.Context, req *pb.ReservationId) (*pb.Status, error) {
	resp, err := r.reser.DeleteReservation(req)
	if err != nil {
		r.Logger.Error("Resirvation deleting error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetReservationsByUserId(ctx context.Context, req *pb.UserId) (*pb.Reservations, error) {
	resp, err := r.reser.GetReservationsByUserId(req)
	if err != nil {
		r.Logger.Error("User id buyicha malumotlarni olishda xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) OrderMeal(ctx context.Context, req *pb.Order) (*pb.Status, error) {
	resp, err := r.reser.OrderMeal(req)
	if err != nil {
		r.Logger.Error("Order service da xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) PayForReservation(ctx context.Context, req *pb.Payment) (*pb.Status, error) {
	resp, err := r.reser.PayForReservation(req)
	if err != nil {
		r.Logger.Error("Payment service da xatolik: %v", err)
		return nil, err
	}
	return resp, nil
}
