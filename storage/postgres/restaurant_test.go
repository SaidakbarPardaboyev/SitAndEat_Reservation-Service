package postgres

import (
	"testing"
	pb "reservation/genproto/reservation"
)

func TestCreateRestaurant(t *testing.T) {
	db,err := ConnectDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := NewRestaurantRepo(db)

	restaurant := &pb.Restuarant{
		Name:        "Test Restaurant",
		Address:     "123 Test St",
		Phone: "1234567890",
		Description: "A place to test",
	}

	status, err := repo.CreateRestaurant(restaurant)
	if err != nil {
		t.Fatalf("Failed to create restaurant: %v", err)
	}

	if !status.True {
		t.Errorf("Expected status true, got false")
	}
}
func TestGetAllRestaurants(t *testing.T) {
	db, err := ConnectDB()

	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := NewRestaurantRepo(db)
	_, err = repo.GetAllRestaurants()
	if err != nil {
		t.Fatalf("Failed to get all restaurants: %v", err)
	}
	// TODO:
}
