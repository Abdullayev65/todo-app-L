CREATE TABLE IF NOT EXISTS users
(
    id            serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS todo_lists
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE IF NOT EXISTS users_lists
(
    id      serial primary key,
    user_id int not null references users (id) on delete cascade,
    list_id int not null references users_lists (id) on delete cascade
);

CREATE TABLE IF NOT EXISTS todo_item
(
    id          serial primary key,
    title       varchar(255) not null,
    description varchar(255) not null,
    done        boolean      not null default false
);

CREATE TABLE IF NOT EXISTS lists_items
(
    id      serial primary key,
    list_id int not null references todo_lists (id) on delete cascade,
    item_id int not null references todo_item (id) on delete cascade
);