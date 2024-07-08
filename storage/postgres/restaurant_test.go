package postgres

import (
	pb "reservation/genproto/reservation"
	"testing"

	_ "github.com/lib/pq"

	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Connect error?")
	}
	defer db.Close()

	return db
}

func TestCreateRestaurant(t *testing.T) {
	db:= Connect()
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
	_, err := repo.GetAllRestaurants(&pb.AllRestuarant{})
	if err != nil {
		t.Fatalf("Failed to get all restaurants: %v", err)
	}

}
