-- Active: 1718919020656@@127.0.0.1@5432@userservice

CREATE TABLE IF NOT EXISTS restaurants_reservation_service (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar,
    address varchar,
    phone_number varchar,
    description varchar,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS reservations_reservation_service (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    restaurant_id uuid,
    reservation_time timestamp default current_timestamp,
    status varchar,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS reservation_orders_reservation_service (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid,
    menu_item_id uuid,
    quantity int,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS menu_reservation_service (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    restaurant_id uuid,
    name varchar,
    description varchar,
    price decimal,
    image bytea,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)