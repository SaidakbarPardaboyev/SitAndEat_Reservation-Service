package main

import (
	"log"
	"net"
	"reservation/config"
	"reservation/genproto/menu"
	"reservation/genproto/menuRedis"
	"reservation/genproto/resirvation"
	"reservation/genproto/restaurant"
	"reservation/service"
	"reservation/storage/postgres"
	"reservation/storage/redis"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().RESERVATION_SERVICE)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	menuService := service.NewMenuService(db)
	reservationService := service.NewReservationService(db)
	restaurantService := service.NewRestaurantService(db)
	redis := redis.NewMenuRedisClient(redis.NewRedisClient())

	service := grpc.NewServer()

	menu.RegisterMenuServer(service, menuService)
	resirvation.RegisterResirvationServer(service, reservationService)
	restaurant.RegisterRestaurantServer(service, restaurantService)
	menuRedis.RegisterMenuServer(service, redis)

	log.Printf("Server is listening on port %s\n", config.Load().RESERVATION_SERVICE)
	if err = service.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
