CREATE TABLE role
(
    id   serial       not null unique primary key,
    name varchar(255) not null
);

CREATE TABLE "user"
(
    id        serial       not null unique primary key,
    full_name varchar(50)  not null,
    login     varchar(255) not null unique,
    password  varchar(255) not null unique,
    role_id   int          not null,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
            REFERENCES role (id)
);

CREATE TABLE taskStatus (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL
);

CREATE TABLE task (
                      id SERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      description TEXT,
                      status_id BIGINT NOT NULL,
                      user_id BIGINT,
                      FOREIGN KEY (status_id) REFERENCES TaskStatus(id),
                      FOREIGN KEY (user_id) REFERENCES "user"(id)
);

INSERT INTO taskStatus (name)
VALUES ('Waiting'), ('Progressing'), ('Completed'), ('Canceled');

INSERT INTO role (name)
VALUES ('admin');
INSERT INTO role (name)
VALUES ('manager');
INSERT INTO role (name)
VALUES ('worker');

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Linus Torvalds', 'admin', '$2a$10$mMj3EqKXqv73Kspkipw/K.jysW3MTwxkKSRSY6qxhdXmsZAnBzxx.', 1);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Mr.Beast', 'manager1', '$2a$10$dAdNLd1fWgoXQv.f3DOHw.GSB./fdzTa8mdkPPf3bKYS6nc69W8bO', 2);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Worker 1', 'worker1', '$2a$10$MfIw7qhpklAsl3/GYZAuc.Vu1zinwYPsPy/OeOxssHGS1N7OXiZyW', 3);
INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Worker 2', 'worker2', '$2a$10$EQBZ4/zj8qPFMGwqZsi55eVsGE1IwyhOwPyhOXzKeqdWjP550O2v2', 3);

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
