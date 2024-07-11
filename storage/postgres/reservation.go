package postgres

import (
	"database/sql"
	pb "reservation/genproto/resirvation"
	"time"
	"github.com/google/uuid"
)

type Reservation struct {
	Db *sql.DB
}

func NewReservationRepo(db *sql.DB) *Reservation {
	return &Reservation{Db: db}
}

func (r *Reservation) CreateReservation(reservation *pb.RequestReservations) (*pb.ReservationId, error) {

	query := `
		INSERT INTO reservations(
			id, user_id, restuarant_id
		) VALUES(
			$1, $2, $3
		)`

	newId := uuid.NewString()
	if _, err := r.Db.Exec(query, newId, reservation.UserId, reservation.RestaurantId); err != nil {
		return &pb.ReservationId{}, err
	}

	return &pb.ReservationId{Id: newId}, nil
}

func (r *Reservation) GetReservationByID(id *pb.ReservationId) (*pb.Reservation, error) {
	reservation := &pb.Reservation{}
	query := `
		SELECT
			id, user_id, restaurant_id, res_time, status, created_at, update_at
		FROM
			reservations
		WHERE
			id = $1 and
			deleted_at is null`
	err := r.Db.QueryRow(query,
		id.Id).Scan(&reservation.Id,
		&reservation.UserId, &reservation.RestuarantId, &reservation.ResTime,
		&reservation.Status, &reservation.CreatedAt, &reservation.UpdateAt,
	)

	return reservation, err
}

func (r *Reservation) GetAllReservation() (*pb.Reservations, error) {
	query := `
			SELECT
				id,	user_id, restaurant_id, res_time, status, created_at, 
				update_at
			FROM
				reservations
			WHERE
				deleted_at is null`
	reservations := []*pb.Reservation{}
	rows, err := r.Db.Query(query)
	if err != nil {
		return &pb.Reservations{Reservations: reservations}, err
	}
	for rows.Next() {
		reservation := &pb.Reservation{}
		err = rows.Scan(&reservation.Id,
			&reservation.UserId, &reservation.RestuarantId, &reservation.ResTime,
			&reservation.Status, &reservation.CreatedAt, &reservation.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return &pb.Reservations{Reservations: reservations}, nil
}

func (r *Reservation) UpdateReservations(res *pb.ReservationUpdate) (*pb.Status, error) {
	query := `
		UPDATE
			reservations
		SET
			restuarant_id = $1
			status = $2,
			update_at = $3
		WHERE
			id = $4 and
			deleted_at is null`

	_, err := r.Db.Exec(query, res.RestuarantId, res.Status, time.Now(), res.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) DeleteReservation(id *pb.ReservationId) (*pb.Status, error) {
	query := `
			UPDATE
				reservations
			SET
				deleted_at = $1
			WHERE
				deleted_at is null and
				id = $2`
	_, err := r.Db.Exec(query, time.Now(), id.Id)
	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) GetReservationsByUserId(id *pb.UserId) (*pb.Reservations, error) {
	query := `
			SELECT
				id, user_id, restaurant_id, res_time, status, created_at, 
				update_at
			FROM
				reservations
			WHERE
				user_id = $1 and
				deleted_at is null
			`

	reservations := []*pb.Reservation{}
	rows, err := r.Db.Query(query,
		id.Id)
	if err != nil {
		return &pb.Reservations{Reservations: reservations}, err
	}
	for rows.Next() {
		reservation := &pb.Reservation{}
		err = rows.Scan(&reservation.Id,
			&reservation.UserId, &reservation.RestuarantId, &reservation.ResTime,
			&reservation.Status, &reservation.CreatedAt, &reservation.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return &pb.Reservations{Reservations: reservations}, nil
}

func (r *Reservation) OrderMeal(order *pb.Order) (*pb.Status, error) {
	query := `
			INSERT INTO reservation_orders (
				reservation_id, menu_item_id, quantity
			) VALUES (
				$1, $2, $3, $4
			)`
	_, err := r.Db.Exec(query, order.ReservatinId, order.MenuItemId,
		order.Quantity)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}

func (r *Reservation) PayForReservation(payment *pb.Payment) (*pb.Status, error) {
	query := `
			INSERT INTO reservation_payments (
				reservation_id,	amount
			) VALUES (
			 	$1, $2
			)`
	_, err := r.Db.Exec(query, payment.ReservationId, payment.Amount)

	if err != nil {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}
