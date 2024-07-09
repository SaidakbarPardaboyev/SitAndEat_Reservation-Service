package postgres

import (
	"database/sql"
	pb "reservation/genproto/resirvation"
	"time"
)

type Reservation struct {
	Db *sql.DB
}

func NewReservationRepo(db *sql.DB) *Reservation {
	return &Reservation{Db: db}
}

func (r *Reservation) CreateReservation(reservation *pb.Reservation) (*pb.Status, error) {

	_, err := r.Db.Exec(`INSERT INTO
							reservations
						(
							id,
							user_id,
							restuarant_id,
							res_time,
							status,
							created_at,
							update_at
						)
						VALUES
						($1,$2,$3,$4,$5,$6,$7)`,
		reservation.Id,
		reservation.UserId,
		reservation.RestuarantId,
		reservation.ResTime,
		reservation.Status,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) GetReservationByID(id *pb.ReservationId) (*pb.Reservation, error) {
	reservation := &pb.Reservation{}
	err := r.Db.QueryRow(`SELECT
							id,
							user_id,
							restaurant_id,
							res_time,
							status,
							created_at,
							update_at
						FROM
							reservations
						WHERE
							id = $1 and
							deleted_at is null`,
		id.Id).Scan(
		&reservation.Id,
		&reservation.UserId,
		&reservation.RestuarantId,
		&reservation.ResTime,
		&reservation.Status,
		&reservation.CreatedAt,
		&reservation.UpdateAt,
	)

	return reservation, err
}

func (r *Reservation) GetAllReservation() (*pb.Reservations, error) {

	reservations := []*pb.Reservation{}
	rows, err := r.Db.Query(`SELECT
								id,
								user_id,
								restaurant_id,
								res_time,
								status,
								created_at,
								update_at
							FROM
								reservations
							WHERE
								deleted_at is null`)
	if err != nil {
		return &pb.Reservations{Reservations: reservations}, err
	}
	for rows.Next() {
		reservation := &pb.Reservation{}
		err = rows.Scan(&reservation.Id,
			&reservation.UserId,
			&reservation.RestuarantId,
			&reservation.ResTime,
			&reservation.Status,
			&reservation.CreatedAt,
			&reservation.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return &pb.Reservations{Reservations: reservations}, nil
}

func (r *Reservation) UpdateReservations(id *pb.Reservation) (*pb.Status, error) {

	_, err := r.Db.Exec(`UPDATE
							reservations
						SET
							status = $1,
							update_at = $2
						WHERE
							id = $3 and
							deleted_at is null`,
		id.Status,
		time.Now(),
		id.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) DeleteReservation(id *pb.ReservationId) (*pb.Status, error) {
	_, err := r.Db.Exec(`UPDATE
							reservations
						SET
							deleted_at = $1
						WHERE
							deleted_at is null and
							id = $2`,
		time.Now(),
		id.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) GetReservationsByUserId(id *pb.UserId) (*pb.Reservations, error) {

	reservations := []*pb.Reservation{}
	rows, err := r.Db.Query(`SELECT
								id,
								user_id,
								restaurant_id,
								res_time,
								status,
								created_at,
								update_at
							FROM
								reservations
							WHERE
								user_id = $1 and
								deleted_at is null`,
		id.Id)
	if err != nil {
		return &pb.Reservations{Reservations: reservations}, err
	}
	for rows.Next() {
		reservation := &pb.Reservation{}
		err = rows.Scan(&reservation.Id,
			&reservation.UserId,
			&reservation.RestuarantId,
			&reservation.ResTime,
			&reservation.Status,
			&reservation.CreatedAt,
			&reservation.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return &pb.Reservations{Reservations: reservations}, nil
}

func (r *Reservation) OrderMeal(order *pb.Order) (*pb.Status, error) {

	_, err := r.Db.Exec(`INSERT INTO
							reservation_orders
								(reservation_id, 
								menu_item_id, 
								quantity) 
							VALUES 

								($1, $2, $3, $4)`,
								order.ReservatinId,
								order.MenuItemId,
								order.Quantity)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) PayForReservation(payment *pb.Payment) (*pb.Status, error) {

	_, err := r.Db.Exec(`INSERT INTO
							reservation_payments
								(reservation_id, 
								amount) 
							VALUES 
								($1, $2)`,
								payment.ReservationId,
								payment.Amount)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}