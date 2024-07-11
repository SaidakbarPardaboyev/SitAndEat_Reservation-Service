package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	pbM "reservation/genproto/menu"
	pb "reservation/genproto/resirvation"
	"reservation/pkg/logger"
	"reservation/storage/postgres"
	"reservation/storage/redis"
)

type ReservationService struct {
	pb.UnimplementedResirvationServer
	Reser     *postgres.Reservation
	MenuRepo  *postgres.Menu
	MenuRedis *redis.MenuRedisClient
	Logger    *slog.Logger
}

func NewReservationService(db *sql.DB) *ReservationService {
	reser := postgres.NewReservationRepo(db)
	menuRepo := postgres.NewMenuRepo(db)
	menu := redis.NewMenuRedisClient(redis.NewRedisClient())
	return &ReservationService{
		Reser:     reser,
		MenuRedis: menu,
		MenuRepo:  menuRepo,
		Logger:    logger.NewLogger(),
	}
}

func (r *ReservationService) Createreservations(ctx context.Context, req *pb.RequestReservations) (*pb.ReservationId, error) {
	resp, err := r.Reser.CreateReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni insert qilishda xatolik: %v", err))
		return nil, err
	}

	meals, err := r.MenuRedis.GetMeals(ctx)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni insert qilishda xatolik: %v", err))
		return nil, err
	}

	uniqueRestaurant := map[string]bool{}
	for ind := range meals.Meals {
		resId, err := r.MenuRepo.GetRestaurantIdByMealId(&pbM.FoodId{Id: meals.Meals[ind].MealId})
		if err != nil {
			r.Logger.Error(fmt.Sprintf("Malumotlarni insert qilishda xatolik: %v", err))
			return nil, err
		}
		if ind == 0 {
			uniqueRestaurant[resId.Id] = true
			continue
		}
		if _, ok := uniqueRestaurant[resId.Id]; !ok {
			r.Reser.DeleteReservation(resp)
			r.Logger.Error(fmt.Sprintf("Faqat bitta restaurantning taomlarini zakas qila olasiz"))
			return nil, err
		}
	}

	for _, meal := range meals.Meals {
		_, err := r.OrderMeal(ctx, &pb.Order{
			MenuItemId:   meal.MealId,
			ReservatinId: resp.Id,
			Quantity:     int32(meal.Quality),
		})
		if err != nil {
			r.Logger.Error(err.Error())
		}
	}

	return resp, nil
}

func (r *ReservationService) GetAllReservations(ctx context.Context, req *pb.FilterField) (*pb.Reservations, error) {
	resp, err := r.Reser.GetAllReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetByIdReservations(ctx context.Context, req *pb.ReservationId) (*pb.Reservation, error) {
	resp, err := r.Reser.GetReservationByID(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Bitta malumotni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) UpdateReservations(ctx context.Context, req *pb.ReservationUpdate) (*pb.Status, error) {
	resp, err := r.Reser.UpdateReservations(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Malumotlarni update qilishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) DeleteReservations(ctx context.Context, req *pb.ReservationId) (*pb.Status, error) {
	resp, err := r.Reser.DeleteReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Resirvation deleting error: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) GetReservationsByUserId(ctx context.Context, req *pb.UserId) (*pb.Reservations, error) {
	resp, err := r.Reser.GetReservationsByUserId(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("User id buyicha malumotlarni olishda xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) OrderMeal(ctx context.Context, req *pb.Order) (*pb.Status, error) {
	resp, err := r.Reser.OrderMeal(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Order service da xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}

func (r *ReservationService) PayForReservation(ctx context.Context, req *pb.Payment) (*pb.Status, error) {
	resp, err := r.Reser.PayForReservation(req)
	if err != nil {
		r.Logger.Error(fmt.Sprintf("Payment service da xatolik: %v", err))
		return nil, err
	}
	return resp, nil
}
