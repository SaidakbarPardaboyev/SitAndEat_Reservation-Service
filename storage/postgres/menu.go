package postgres

import (
	"database/sql"
	"time"

	menu "reservation/genproto/menu"
)

type Menu struct {
	Db *sql.DB
}

func NewMenuRepo(db *sql.DB) *Menu {
	return &Menu{Db: db}
}

func (m *Menu) CreateFood(food *menu.CreateF) (*menu.Status, error) {
	query := `
	INSERT INTO menu(
		restaurant_id, name, description, price, image
	) VALUES (
		$1, $2, $3, $4, $5
	)`
	_, err := m.Db.Exec(query, food.RestuarantId, food.Name, food.Description,
		food.Price, food.Image)
	if err != nil {
		return nil, err
	}
	return &menu.Status{Status: true}, nil
}

func (m *Menu) GetAllFoods() (*menu.Foods, error) {
	query := `
		SELECT
			id, restaurant_id, name, description, price, image, created_at, 
			update_at
		FROM
			menu
		WHERE
			deleted_at is null
		ORDER BY
			name`
	rows, err := m.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fuds := &menu.Foods{}
	for rows.Next() {
		fud := &menu.Food{}
		err = rows.Scan(
			&fud.Id, &fud.RestuarantId, &fud.Name, &fud.Description,
			&fud.Price, &fud.Image, &fud.CreatedAt, &fud.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		fuds.Foods = append(fuds.Foods, fud)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return fuds, nil
}

func (m *Menu) GetFood(food *menu.FoodId) (*menu.Food, error) {
	query := `
		SELECT
			id, restaurant_id, name, description, price, image, created_at, 
			update_at
		FROM
			menu
		WHERE
			deleted_at is null AND
			id = $1
	`
	row := m.Db.QueryRow(query, food.Id)
	fud := &menu.Food{}
	err := row.Scan(
		&fud.Id, &fud.RestuarantId, &fud.Name, &fud.Description,
		&fud.Price, &fud.Image, &fud.CreatedAt, &fud.UpdateAt,
	)
	if err != nil {
		return nil, err
	}
	return fud, nil
}

func (m *Menu) UpdateFood(food *menu.UpdateF) (*menu.Status, error) {
	query := `
		UPDATE
			menu
		SET
			name = $1, 
			description = $2, 
			price = $3, 
			image = $4
		WHERE
			deleted_at is null AND
			id = $5
		`
	_, err := m.Db.Exec(query, food.Name, food.Description, food.Price,
		food.Image, food.Id)
	if err != nil {
		return nil, err
	}
	return &menu.Status{Status: true}, nil
}

func (m *Menu) DeleteFood(food *menu.FoodId) (*menu.Status, error) {
	query := `
		UPDATE
			menu
		SET
			deleted_at = $1
		WHERE
			deleted_at is null AND
			id = $2
		`

	_, err := m.Db.Exec(query, time.Now(), food.Id)
	if err != nil {
		return nil, err
	}
	return &menu.Status{Status: true}, nil
}

func (m *Menu) GetRestaurantIdByMealId(food *menu.FoodId) (*menu.FoodId, error) {
	query := `
		select
			restaurant_id
		from 
			menu
		where
			id = $1 and
			deleted_at is null
	`
	var id string
	err := m.Db.QueryRow(query, food.Id).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &menu.FoodId{Id: id}, nil
}
