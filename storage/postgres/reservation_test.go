package postgres

import (
	"reflect"
	pb "reservation/genproto/resirvation"
	"testing"
)

func TestCreateReservation(t *testing.T){
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	reservation := &pb.RequestReservations{
		UserId: "",
		RestaurantId: "",
	}

	status, err := repo.CreateReservation(reservation)
	if err != nil{
		t.Error(err)
	}

	if !status.Status{
		t.Error("Service xato ishladi.")
	}
}

func TestGetReservationByID(t *testing.T){
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	resId := &pb.ReservationId{
		Id: "",
	}

	reservation, err := repo.GetReservationByID(resId)
	if err != nil{
		t.Error(err)
	}

	reservationDb := &pb.Reservation{
		Id: "",
		UserId: "",
		RestuarantId: "",
		ResTime: "",
		CreatedAt: "",
		UpdateAt: "",
		Status: "",
	}

	if !reflect.DeepEqual(reservation, reservationDb){
		t.Error("Ma'lumotlar mos kelmadi.")
	}
}


func TestGetAllReservation(t *testing.T){
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	_, err := repo.GetAllReservation()
	if err != nil{
		t.Error(err)
	}
}

func TestUpdateReservations(t *testing.T){
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	reservation := &pb.ReservationUpdate{
		
	}
}