create table users 
(
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) unique
);

create table todo_lists 
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255) 
);

create table users_lists
(
    id serial primary key,
    user_id int references users(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

create table todo_item
(
    id serial primary key,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null default false
);

create table lists_items
(
    id serial primary key,
    item_id int references todo_item(id) on delete cascade not null,
    list_id int references todo_lists(id) on delete cascade not null
);

