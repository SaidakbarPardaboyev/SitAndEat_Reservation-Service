package postgres

import "database/sql"

type NewRestaurant struct{
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB)*NewRestaurant{
	return &NewRestaurant{Db: db}
}
