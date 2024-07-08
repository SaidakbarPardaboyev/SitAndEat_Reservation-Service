package postgres

import (
	"database/sql"
	pb "reservation/genproto/reservation"
)

type NewRestaurant struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *NewRestaurant {
	return &NewRestaurant{Db: db}
}

func (r *NewRestaurant) CreateRestaurant(restaurant *pb.Restuarant) (*pb.Status, error) {
	_, err := r.Db.Exec("INSERT INTO restaurants (name, address, phone_number, description) VALUES ($1, $2, $3, $4)",
		restaurant.Name, restaurant.Address, restaurant.Phone, restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &pb.Status{True: true}, nil
}

func (r *NewRestaurant) GetAllRestaurants() (*pb.Restuanants, error) {
	rows, err := r.Db.Query("SELECT * FROM restaurants")
	if err != nil {
		return nil, err
	}
	var restaurants []*pb.GetRes
	for rows.Next() {
		restaurant := &pb.GetRes{}
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address, &restaurant.Phone, &restaurant.Description, &restaurant.CreatedAt, &restaurant.UpdatedAt)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	return &pb.Restuanants{Restuanants: restaurants}, nil
}