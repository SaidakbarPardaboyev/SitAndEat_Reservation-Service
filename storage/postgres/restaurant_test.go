package postgres

import (
	"database/sql"
	"log"
	"testing"

	pb "reservation/genproto/restaurant"

	_ "github.com/lib/pq"
)

// Connect initializes the database connection and logs any connection errors.
func Connect() *sql.DB {
	db, err := ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func TestCreateRestaurant(t *testing.T) {
	db := Connect()
	defer db.Close()
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
	defer db.Close()

	repo := NewRestaurantRepo(db)

	restaurants, err := repo.GetAllRestaurants(&pb.FilterField{})
	if err != nil {
		t.Fatalf("Failed to get all restaurants: %v", err)
	}

	if len(restaurants.Restuanants) == 0 {
		t.Fatalf("Expected at least one restaurant, got none")
	}
}

func TestGetRestaurant(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewRestaurantRepo(db)

	restaurantId := &pb.RestuanantId{Id: "41fcc8a7-1a85-49aa-8578-476d18eb7f4b"}
	_, err := repo.GetRestaurant(restaurantId)
	if err != nil {
		t.Fatalf("Failed to get restaurant: %v", err)
	}
}

func TestUpdateRestuarant(t *testing.T) {
	db := Connect()
	defer db.Close()

	repo := NewRestaurantRepo(db)

	updateRestaurant := &pb.RestuarantUpdate{
		Id:          "b75236c6-69b1-4f40-b53d-e9667280c208",
		Name:        "Updated Name",
		Address:     "Updated Address",
		Phone:       "998 99 4558545",
		Description: "Updated description",
	}

	status, err := repo.UpdateRestaurant(updateRestaurant)
	if err != nil {
		t.Fatalf("Failed to update restaurant: %v", err)
	}

	if !status.Status {
		t.Errorf("Expected status true, got false")
	}
}

func TestDeleteRestaurant(t *testing.T) {
	db := Connect()
	defer db.Close()
	repo := NewRestaurantRepo(db)

	restaurantId := &pb.RestuanantId{
		Id: "b75236c6-69b1-4f40-b53d-e9667280c208",
	}

	status, err := repo.DeleteRestaurant(restaurantId)
	if err != nil {
		t.Fatalf("Failed to delete restaurant: %v", err)
	}

	if !status.Status {
		t.Errorf("Expected status true, got false")
	}
}
