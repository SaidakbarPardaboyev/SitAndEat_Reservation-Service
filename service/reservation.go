package service

import (
	"context"
	"database/sql"
	"log"
	pb "reservation/genproto/reservation"
	"reservation/storage/postgres"
)
type ReservationService struct{
	pb.UnimplementedResirvationServer
	db *sql.DB
	reser *postgres.NewRestaurant
}

func NewRestaurantService(db *sql.DB,reser *postgres.NewRestaurant) *ReservationService{
	return &ReservationService{db: db,reser: reser}
}

func (s *ReservationService) CreateRestaurant(ctx context.Context,req *pb.Restuarant) (*pb.Status,error){
	resp,err:=s.reser.CreateRestaurant(req)
	if err!=nil{
		log.Fatal("Data inserting?",err)
		return nil,err
	}
	return resp,nil
}

func (s *ReservationService) GetAllRestaurants(ctx context.Context,req *pb.AllRestuarant)(*pb.Restuanants,error){
	resp,err:=s.reser.GetAllRestaurants(req)
	if err!=nil{
		log.Fatal("GetAll error?",err)
		return nil,err
	}
	return resp,nil
}

func (s *ReservationService) GetRestuarant(ctx context.Context,req *pb.RestuanantId)(*pb.GetRes,error){
	resp,err:=s.reser.GetRestuarant(req)
	if err!=nil{
		log.Fatal("GetRestaurant error?")
		return nil,err
	}
	return resp,nil
}

func (s *ReservationService) UpdateRestuarant(ctx context.Context,req *pb.GetRes)(*pb.Status,error){
	resp,err:=s.reser.UpdateRestuarant(req)
	if err!=nil{
		log.Fatal("UpdateRestaurant error?")
		return nil,err
	}
	return resp,nil
}

func (s *ReservationService) DeleteRestuarant(ctx context.Context,req *pb.RestuanantId)(*pb.Status,error){
	resp,err:=s.reser.DeleteRestuarant(req)
	if err!=nil{
		log.Fatal("DeleteRestaurant error?",err)
		return nil,err
	}
	return resp,nil
}