package redis

import (
	"context"
	pb "reservation/genproto/menuRedis"
	"testing"
)

func Test(t *testing.T) {
	redisclient := NewRedisClient()
	client := NewMenuRedisClient(redisclient)
	status, err := client.CretaeMeal(context.Background(), &pb.MealCreate{
		MealId:  "alik",
		Quality: 32,
	})
	if err != nil {
		t.Error(err)
	}
	if !status.Status {
		t.Error("not created")
	}
}
func TestGetMeals(t *testing.T) {
	redisclient := NewRedisClient()
	client := NewMenuRedisClient(redisclient)
	_, err := client.GetMeals(context.Background(), &pb.Void{})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateMeal(t *testing.T) {
	redisclient := NewRedisClient()
	client := NewMenuRedisClient(redisclient)
	status, err := client.UpdateMeal(context.Background(), &pb.MealCreate{
		MealId:  "alik",
		Quality: 35,
	})
	if err != nil {
		t.Error(err)
	}
	if !status.Status {
		t.Error("not created")
	}
}

func TestDeleteMeal(t *testing.T) {
	redisclient := NewRedisClient()
	client := NewMenuRedisClient(redisclient)
	status, err := client.DeleteMeal(context.Background(), &pb.MealDelete{
		MealId: "alik",
	})
	if err != nil {
		t.Error(err)
	}
	if !status.Status {
		t.Error("not created")
	}
}
