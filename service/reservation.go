package service

import (
	"context"
	"database/sql"
	"log"
	pb "reservation/genproto/resirvation"
	"reservation/storage/postgres"
)
type ReservationService struct {
	pb.UnimplementedResirvationServer
	db *sql.DB
	reser *postgres.Reservation
}

func NewReservationService(db *sql.DB,reser *postgres.Reservation) *ReservationService{
	return &ReservationService{db: db,reser: reser}
}

func (r *ReservationService) Createreservations(ctx context.Context,req *pb.Reservation)(*pb.Status,error){
	resp,err:=r.reser.CreateReservation(req)
	if err != nil {
		log.Fatalf("Malumotlarni insert qilishda xatolik: %v",err)
		return nil, err
	}

	return resp,nil
}

func (r *ReservationService) GetAllReservations(ctx context.Context,req *pb.Void)(*pb.Reservations,error){
	resp,err:=r.reser.GetAllReservation()
	if err!=nil{
		log.Fatalf("Malumotlarni olishda xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) GetByIdReservations(ctx context.Context,req *pb.ReservationId)(*pb.Reservation,error){
	resp,err:=r.reser.GetReservationByID(req)
	if err!=nil{
		log.Fatalf("Bitta malumotni olishda xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) UpdateReservations(ctx context.Context,req *pb.ReservationUpdate)(*pb.Status,error){
	resp,err:=r.reser.UpdateReservations(req)
	if err!=nil{
		log.Fatalf("Malumotlarni update qilishda xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) DeleteReservations(ctx context.Context,req *pb.ReservationId)(*pb.Status,error){
	resp,err:=r.reser.DeleteReservation(req)
	if err!=nil{
		log.Fatalf("Resirvation deleting error: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) GetReservationsByUserId(ctx context.Context,req *pb.UserId)(*pb.Reservations,error){
	resp,err:=r.reser.GetReservationsByUserId(req)
	if err!=nil{
		log.Fatalf("User id buyicha malumotlarni olishda xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) OrderMeal(ctx context.Context,req *pb.Order)(*pb.Status,error){
	resp,err:=r.reser.OrderMeal(req)
	if err!=nil{
		log.Fatalf("Order service da xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}

func (r *ReservationService) PayForReservation(ctx context.Context,req *pb.Payment)(*pb.Status,error){
	resp,err:=r.reser.PayForReservation(req)
	if err!=nil{
		log.Fatalf("Payment service da xatolik: %v",err)
		return nil,err
	}
	return resp,nil
}