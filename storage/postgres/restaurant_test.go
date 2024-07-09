package postgres

import (
	"database/sql"
	"log"
	pb "reservation/genproto/reservation"
	"testing"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Connect error?")
	}

	return db
}

func TestCreateRestaurant(t *testing.T) {
	db := Connect()
	repo := NewRestaurantRepo(db)

	restaurant := &pb.Restuarant{
		Name:        "Test Restaurant",
		Address:     "123 Test St",
		Phone:       "1234567890",
		Description: "A place to test",
	}

	status, err := repo.CreateRestaurant(restaurant)
	if err != nil {
		t.Fatalf("Failed to create restaurant: %v", err)
	}

	if !status.Status {
		t.Errorf("Expected status true, got false")
	}

}

func TestGetAllRestaurants(t *testing.T) {
	db := Connect()

	repo := NewRestaurantRepo(db)
	restaurant, err := repo.GetAllRestaurants(&pb.AllRestuarant{})
	if err != nil {
		t.Fatalf("Failed to get all restaurants: %v", err)
	}

	if len(restaurant.Restuanants) == 0 {
		t.Fatalf("Kamida bitta restaurant bo'lishi kerak edi lekin topilmadi?")
	}
}

func TestGetRestuarant(t *testing.T){
	db := Connect()

	repo:=NewRestaurantRepo(db)

	restyaurantId:=&pb.RestuanantId{RestuarantId: "41fcc8a7-1a85-49aa-8578-476d18eb7f4b"}
	_,err:=repo.GetRestuarant(restyaurantId)
	if err!=nil{
		t.Fatalf("Restaurant olishda xatolik berdi: %v",err)
	}
}

func TestUpdateRestuarant(t *testing.T){
	db:=Connect()
	repo:=NewRestaurantRepo(db)

	updaterestaurant:=&pb.GetRes{
		Id: "b75236c6-69b1-4f40-b53d-e9667280c208",
		Name: "nimadur",
		Address: "Qayerdir",
		Phone: "998 99 4558545",
		Description: "description",
	}
	_,err:=repo.UpdateRestuarant(updaterestaurant)
	if err!=nil{
		t.Fatalf("Update testing error: %v",err)
	}
}

func TestDeleteRestaurant(t *testing.T){
	db:=Connect()

	repo:=NewRestaurantRepo(db)
	restaurantId:=&pb.RestuanantId{
		RestuarantId: "b75236c6-69b1-4f40-b53d-e9667280c208",
	}

	_,err:=repo.DeleteRestuarant(restaurantId)
	if err!=nil{
		t.Fatalf("Error delete testing restaurant: %v",err)
	}
}