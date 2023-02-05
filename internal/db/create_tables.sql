-- Create users table
create table if not exists users(
    id serial primary key,
    username varchar(20),
    password varchar(255),
    email varchar(255) unique,
    phone varchar(10)
);

-- Create accounts table
create table if not exists accounts(
    id serial primary key,
    name varchar(20),
    user_id int,

    constraint FK_accounts_users FOREIGN KEY(user_id) REFERENCES users(id)
);
