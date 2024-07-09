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
	_, err := repo.CreateFood(food)
	if err != nil {
		t.Fatalf("Create testing error: %v", err)
	}
}
