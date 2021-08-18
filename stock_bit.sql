-- Soal No 1

create table if not exists users
(
    id       serial primary key,
    username varchar(255),
    parent   int
);

insert into users
values ('Ali', 2);
insert into users
values ('Budi', 0);
insert into users
values ('Cecep', 1);

select u2.id, u2.username, (select u1.username from users u1 where u1.id = u2.parent) as parent_username
from users u2;


-- Soal Nomor 2

create table if not exists movies
(
    id      serial primary key,
    title   varchar(55) unique,
    year    varchar(55),
    imdb_id varchar(55),
    type    varchar(55),
    poster  varchar(255)
);

create table if not exists movies_detail
(
    id       serial primary key,
    title    varchar(55) unique,
    year     varchar(55),
    rated    varchar(55),
    released varchar(55),
    runtime  varchar(55),
    genre    varchar(55),
    director varchar(55),
    writer   varchar(55),
    actors   text,
    plot     text,
    language varchar(55),
    country  varchar(255),
    awards   text,
    poster   varchar(255)
);
