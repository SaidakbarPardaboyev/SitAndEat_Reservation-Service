package postgres

import (
	"database/sql"
	pb "reservation/genproto/resirvation"
	"time"
)

type NewRestaurant struct{
	Db *sql.DB
}

func NewRestaurantRepo(db *sql.DB)*NewRestaurant{
	return &NewRestaurant{Db: db}
}


func(R *NewRestaurant) GetRestuarant(id *pb.RestuanantId)(*pb.GetRes, error){
	restaurant := &pb.GetRes{}
	err := R.Db.QueryRow(`SELECT
							id, 
							name, 
							address, 
							phone_number,
							description, 
							created_at,
							update_at
						FROM 
							restaurants_reservation_service
						WHERE
							id = $1 and
							deleted_at is not null`, 
						id.RestuarantId).Scan(
							&restaurant.Id,
							&restaurant.Name,
							&restaurant.Address,
							&restaurant.Phone,
							&restaurant.Description,
							&restaurant.CreatedAt,
							&restaurant.UpdatedAt,
						)
	return restaurant, err
}


func(R *NewRestaurant) UpdateRestuarant(restuarant *pb.GetRes)(*pb.Status, error){
	_, err := R.Db.Exec(`UPDATE 
							restaurants_reservation_service
						SET
							name = $1,
							address = $2,
							phone_number = $3,
							description = $4,
							update_at = $5
						WHERE 
							id = $6 and
							deleted_at is not null`, 
							restuarant.Name,
							restuarant.Address,
							restuarant.Phone,
							restuarant.Description,
							time.Now(),
							restuarant.Id)	
	if err != nil{
		return &pb.Status{Status: false}, err
	}	
	return &pb.Status{Status: true}, nil
}

func(R *NewRestaurant) DeleteRestuarant(id *pb.RestuanantId)(*pb.Status, error){
	_, err := R.Db.Exec(`UPDATE
							restaurants_reservation_service
						SET
							deleted_at = $1
						WHERE
							deleted_at is not null and
							id = $2`, 
							time.Now(),
							id.RestuarantId)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}