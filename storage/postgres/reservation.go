package postgres

import "database/sql"

type NewReservation struct{
	Db *sql.DB
}

func NewReservationRepo(db *sql.DB)*NewReservation{
	return &NewReservation{Db: db}
}