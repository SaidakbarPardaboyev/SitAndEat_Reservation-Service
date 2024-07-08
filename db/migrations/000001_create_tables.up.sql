
CREATE TABLE IF NOT EXISTS restaurants(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar not null,
    address varchar not null,
    phone_number varchar not null,
    description varchar,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS reservations(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid,
    restaurant_id uuid,
    reservation_time timestamp default current_timestamp,
    status varchar,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS reservation_orders(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id uuid not null,
    menu_item_id uuid not null,
    quantity int not null,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)

CREATE TABLE IF NOT EXISTS menu(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    restaurant_id uuid not null,
    name varchar not null,
    description varchar,
    price decimal not null,
    image bytea,
    created_at timestamp default current_timestamp,
    update_at timestamp default current_timestamp,
    deleted_at timestamp
)