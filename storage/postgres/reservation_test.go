package postgres

import (
	"reflect"
	pb "reservation/genproto/resirvation"
	"testing"
)

func TestCreateReservation(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	reservation := &pb.RequestReservations{
		UserId:       "",
		RestaurantId: "",
	}

	_, err := repo.CreateReservation(reservation)
	if err != nil {
		t.Error(err)
	}
}

func TestGetReservationByID(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	resId := &pb.ReservationId{
		Id: "",
	}

	reservation, err := repo.GetReservationByID(resId)
	if err != nil {
		t.Error(err)
	}

	reservationDb := &pb.Reservation{
		Id:           "",
		UserId:       "",
		RestuarantId: "",
		ResTime:      "",
		CreatedAt:    "",
		UpdateAt:     "",
		Status:       "",
	}

	if !reflect.DeepEqual(reservation, reservationDb) {
		t.Errorf("GetReservationByID methodda Ma'lumotlar mos kelmadi.")
	}
}

func TestGetAllReservation(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	_, err := repo.GetAllReservation(&pb.FilterField{})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateReservations(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	reservation := &pb.ReservationUpdate{
		Id:           "",
		RestuarantId: "",
		Status:       "",
	}

	status, err := repo.UpdateReservations(reservation)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Errorf("Ma'lumot yangilanmadi")
	}
}

func TestDeleteReservation(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	id := &pb.ReservationId{
		Id: "",
	}

	status, err := repo.DeleteReservation(id)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Errorf("Ma'lumot o'chirilmadi.")
	}
}

func TestGetReservationsByUserId(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	id := &pb.UserId{}

	_, err := repo.GetReservationsByUserId(id)
	if err != nil {
		t.Error(err)
	}
}

func TestOrderMeal(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	order := &pb.Order{
		ReservatinId: "",
		MenuItemId:   "",
		Quantity:     0,
	}

	status, err := repo.OrderMeal(order)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Errorf("Buyurtma qabul qilinmadi.")
	}
}

func TestPayForReservation(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewReservationRepo(db)

	payment := &pb.Payment{
		ReservationId: "",
		Amount:        0.0,
	}

	status, err := repo.PayForReservation(payment)
	if err != nil {
		t.Error(err)
	}

	if !status.Status {
		t.Errorf("To'lov qilishda xatolik.")
	}
}
