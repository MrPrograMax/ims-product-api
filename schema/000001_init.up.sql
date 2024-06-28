create table category
(
    id     bigserial primary key,
    "name" varchar(255) not null unique
);

create table location
(
    id    bigserial primary key,
    "row" varchar(32) not null,
    place varchar(32) not null
);

create table product_status
(
    id     bigserial primary key,
    "name" varchar(255) not null unique
);

create table product
(
    id          bigserial primary key,
    "name"      varchar(255) not null,
    quantity    bigint       not null,
    description text,
    category_id bigint       not null references category (id),
    location_id bigint       not null references location (id),
    status_id   bigint       not null references product_status (id)
);

create table supply
(
    id       bigserial primary key,
    datetime timestamp not null
);

create table supply_item
(
    id         bigserial primary key,
    supply_id  bigint not null references supply (id),
    product_id bigint not null references product (id),
    quantity   bigint not null
);

create table "order"
(
    id       bigserial primary key,
    datetime timestamp not null
);

create table order_item
(
    id         bigserial primary key,
    order_id  bigint not null references "order" (id),
    product_id bigint not null references product (id),
    quantity   bigint not null
);

insert into product_status(name)
values ('active'),
       ('inactive');

insert into category(name)
values ('Продукты'),
       ('Одежда'),
       ('Техника');

