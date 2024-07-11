package postgres

import (
	"database/sql"
	pb "reservation/genproto/restaurant"
	"strconv"
	"strings"
	"time"
)

type NewRestaurant struct {
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB) *NewRestaurant {
	return &NewRestaurant{Db: db}
}

func (r *NewRestaurant) GetRestaurant(id *pb.RestuanantId) (*pb.GetRes, error) {
	restaurant := &pb.GetRes{}
	query := `
		SELECT
			id, name, address, phone_number, description, created_at, updated_at
		FROM 
			restaurants
		WHERE
			id = $1 AND
			deleted_at IS NULL`
	err := r.Db.QueryRow(query, id.Id).Scan(&restaurant.Id,
		&restaurant.Name, &restaurant.Address, &restaurant.Phone,
		&restaurant.Description, &restaurant.CreatedAt, &restaurant.UpdatedAt,
	)
	return restaurant, err
}

func (r *NewRestaurant) UpdateRestaurant(restaurant *pb.RestuarantUpdate) (*pb.Status, error) {
	query := `
		UPDATE
			restaurants
		SET
			name = $1,
			address = $2,
			phone_number = $3,
			description = $4,
			updated_at = $5
		WHERE 
			id = $6 AND
			deleted_at IS NULL`
	_, err := r.Db.Exec(query,
		restaurant.Name, restaurant.Address, restaurant.Phone,
		restaurant.Description, time.Now(), restaurant.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *NewRestaurant) DeleteRestaurant(id *pb.RestuanantId) (*pb.Status, error) {
	query := `
		UPDATE 
			restaurants
		SET
			deleted_at = $1
		WHERE
			id = $2 AND
			deleted_at IS NULL`
	_, err := r.Db.Exec(query, time.Now(), id.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *NewRestaurant) CreateRestaurant(restaurant *pb.Restuarant) (*pb.Status, error) {
	query := `
		INSERT INTO restaurants(
				name, address, phone_number, description
		) VALUES (
			$1, $2, $3, $4
		)`
	_, err := r.Db.Exec(query, restaurant.Name, restaurant.Address,
		restaurant.Phone, restaurant.Description)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *NewRestaurant) GetAllRestaurants(field *pb.FilterField) (*pb.Restuanants, error) {

	query := `
		  SELECT 
			* 
		  FROM 
			Restaurants
		  WHERE 
			deleted_at is null`
	param := []string{}
	arr := []interface{}{}

	if len(field.Name) > 0 {
		query += " and name = :name"
		param = append(param, ":name")
		arr = append(arr, field.Name)
	}

	if len(field.Address) > 0 {
		query += " and address = :address"
		param = append(param, ":address")
		arr = append(arr, field.Address)
	}

	if len(field.CreatedAt) > 0 {
		data := strings.Split(field.CreatedAt, "-")
		query += " and created_at BETWEEN :created_at1 and :created_at2"
		param = append(param, ":created_at1", ":created_at2")
		arr = append(arr, data[0], data[1])
	}

	if len(field.Limit) > 0 {
		query += " limit " + field.Limit
	}

	if len(field.Offset) > 0 {
		query += " offset " + field.Offset
	}

	for k, v := range param {
		query = strings.Replace(query, v, "$"+strconv.Itoa(k+1), 1)
	}

	rows, err := r.Db.Query(query, arr...)
	if err != nil {
		return &pb.Restuanants{}, err
	}
	defer rows.Close()

	var restaurants []*pb.GetRes
	for rows.Next() {
		restaurant := &pb.GetRes{}
		err = rows.Scan(&restaurant.Id, &restaurant.Name, &restaurant.Address,
			&restaurant.Phone, &restaurant.Description, &restaurant.CreatedAt,
			&restaurant.UpdatedAt)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, restaurant)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &pb.Restuanants{Restuanants: restaurants}, nil
}
