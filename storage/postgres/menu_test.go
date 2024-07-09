package postgres

import (
	pb "reservation/genproto/menu"
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateFood(t *testing.T) {
	db := Connect()

	repo := NewMenuRepo(db)
	food := &pb.CreateF{
		RestuarantId: "b75236c6-69b1-4f40-b53d-e9667280c208",
		Name:         "Nimadur",
		Description:  "description",
		Price:        10000,
		Image:        []byte("image"),
	}
	status, err := repo.CreateFood(food)
	if err != nil {
		t.Fatalf("Create testing error: %v", err)
	}

	if !status.Status {
		t.Errorf("Expected status true, got false")
	}

	db.Close()

	t.Log("Test passed")
}

func TestGetAllFoods(t *testing.T) {
	db := Connect()

	repo := NewMenuRepo(db)
	foods, err := repo.GetAllFoods(&pb.Void{})
	if err != nil {
		t.Fatalf("GetAllFoods testing error: %v", err)
	}

	if len(foods.Foods) == 0 {
		t.Fatalf("Expected at least one food, got none")
	}
}

func TestGetFood(t *testing.T) {
	db := Connect()
	repo := NewMenuRepo(db)
	foodId := &pb.FoodId{Id: "ff671b68-8497-4b42-b26b-a213a16abb7e"}
	food, err := repo.GetFood(foodId)
	if err != nil {
		t.Fatalf("GetFood testing error: %v", err)
	}

	if food.Name == "" {
		t.Fatalf("Expected food name, got none")
	}
	if food.Description == "" {
		t.Fatalf("Expected food description, got none")
	}
	if food.Price == 0 {
		t.Fatalf("Expected food price, got none")
	}
	if food.Image == nil {
		t.Fatalf("Expected food image, got none")
	}
	db.Close()
	t.Log("Test passed")
}

func TestUpdateFood(t *testing.T) {
	db := Connect()
	repo := NewMenuRepo(db)
	food := &pb.UpdateF{
		Id:          "b75236c6-69b1-4f40-b53d-e9667280c208",
		Name:        "Nimadur-test",
		Description: "description",
		Price:       106,
		Image:       []byte("image"),
	}
	status, err := repo.UpdateFood(food)
	if err != nil {
		t.Fatalf("UpdateFood testing error: %v", err)
	}
	if !status.Status {
		t.Errorf("Expected status true, got false")
	}
	db.Close()
	t.Log("Test passed")
}

func TestDeleteFood(t *testing.T) {
	db := Connect()
	repo := NewMenuRepo(db)
	foodId := &pb.FoodId{Id: "b75236c6-69b1-4f40-b53d-e9667280c208"}
	status, err := repo.DeleteFood(foodId)
	if err != nil {
		t.Fatalf("DeleteFood testing error: %v", err)
	}
	if !status.Status {
		t.Errorf("Expected status true, got false")
	}
	db.Close()
	t.Log("Test passed")
}
