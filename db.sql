CREATE TABLE category(
    id uuid primary key,
    namec varchar(10),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

CREATE TABLE products(
    id uuid ,
    name_ varchar(10),
    category_id uuid references category(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp ,
    price int,

);