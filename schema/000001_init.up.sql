create table category
(
    id   serial primary key,
    "name" varchar(255) not null unique
);

create table location
(
    id    serial primary key,
    "row" varchar(32) not null,
    place varchar(32) not null
);

create table product_status
(
    id   serial primary key,
    "name" varchar(255) not null unique
);

create table product
(
    id          serial primary key,
    "name"        varchar(255) not null,
    quantity    bigint       not null,
    description text,
    category_id bigint       not null references category (id),
    location_id bigint       not null references location (id),
    status_id   bigint       not null references product_status (id)
);

insert into product_status(name)
values ('active'),
       ('inactive');

insert into category(name)
values ('Продукты'),
       ('Одежда'),
       ('Техника');
