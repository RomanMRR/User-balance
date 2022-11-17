CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE wallet
(
    id serial not null unique,
    amount float DEFAULT 0,
    user_id int references users (id) on delete cascade not null 
);

CREATE TABLE transaction
(
    user_id int references users (id) on delete cascade not null,
    service_id int not null,
    order_id int not null,
    sum float not null,
    date timestamp not null
);