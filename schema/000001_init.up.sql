CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    balance float DEFAULT 0,
    reserved_balance float DEFAULT 0
);

CREATE TABLE transaction
(
    user_id int references users (id) on delete cascade not null,
    service_id int not null,
    order_id int not null,
    sum float not null,
    date timestamp not null
);