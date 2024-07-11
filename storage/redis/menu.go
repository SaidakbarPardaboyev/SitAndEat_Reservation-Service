package redis

import (
	pb "reservation/genproto/menuRedis"
	"strconv"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type MenuRedisClient struct {
	Redis *redis.Client
}

func NewMenuRedisClient(redis *redis.Client) *MenuRedisClient {
	return &MenuRedisClient{
		Redis: redis,
	}
}

func (m *MenuRedisClient) CretaeMeal(ctx context.Context, meal *pb.MealCreate) (*pb.Status, error) {
	err := m.Redis.HSet(ctx, "orders", meal.MealId, meal.Quality).Err()
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (m *MenuRedisClient) GetMeals(ctx context.Context) (*pb.Meals, error) {
	orders, err := m.Redis.HGetAll(ctx, "orders").Result()
	if err != nil {
		return nil, err
	}

	res := pb.Meals{}
	for k, v := range orders {
		quantity, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		res.Meals = append(res.Meals, &pb.MealCreate{
			MealId:  k,
			Quality: int64(quantity),
		})
	}

	return &res, nil
}

func (m *MenuRedisClient) UpdateMeal(ctx context.Context, meal *pb.MealCreate) (*pb.Status, error) {
	err := m.Redis.HSet(ctx, "orders", meal.MealId, meal.Quality).Err()
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (m *MenuRedisClient) DeleteMeal(ctx context.Context, meal *pb.MealDelete) (*pb.Status, error) {
	err := m.Redis.HDel(ctx, "orders", meal.MealId).Err()
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}
