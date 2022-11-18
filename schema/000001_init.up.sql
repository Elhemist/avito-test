CREATE TABLE IF NOT EXISTS users
(
    id      serial              primary key,
    balance numeric             not null,
    reserve numeric DEFAULT 0   not null 
);

CREATE TABLE IF NOT EXISTS services
(
    id      serial          primary key,
    name    varchar(255)    not null,
    price   numeric         not null
);

CREATE TABLE IF NOT EXISTS orders
(
    id          serial      primary key,
    serviceId   INTEGER,
    userId      INTEGER,
    date        timestamp   default CURRENT_TIMESTAMP,
    status      boolean     default false
);