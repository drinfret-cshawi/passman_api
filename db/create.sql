create table users
(
    id       int generated always as identity primary key,
    username varchar(20) not null unique check ( length(username) >= 1 ),
    fullname text,
    password text        not null check ( length(password) >= 8 ),
    email    text
);

insert into users (username, fullname, password, email)
values ('denis', 'Denis', '12345678', 'denis@example.com'),
       ('alice', 'Alice', '12123232', null);

create table passwords
(
    id       int generated always as identity primary key,
    user_id  int  not null references users (id),
    site     text not null,
    login text,
    password text not null,
    unique (user_id, site, login)
);

insert into passwords(user_id, site, login, password)
VALUES (1, 'google.com', 'denis', '12345678'),
       (1, 'facebook.com', 'denis', '87654321');

create table user_data
(
    id      int generated always as identity primary key,
    user_id int  not null references users (id),
    key     text not null,
    data    text not null
);

