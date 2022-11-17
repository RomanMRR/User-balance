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
    amount NUMERIC(15, 2) DEFAULT 0 CHECK (amount >= 0),
    user_id int references users (id) on delete cascade not null 
);



CREATE TABLE reserve_wallet
(
    id serial not null unique,
    amount NUMERIC(15, 2)  DEFAULT 0 CHECK (amount >= 0),
    user_id int references users (id) on delete cascade not null 
);