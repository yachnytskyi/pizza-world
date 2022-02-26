CREATE TABLE Users 
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE Recipes 
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE Recipe_Users 
(
    id serial not null unique,
    user_id int references Users(id) on delete cascade not null,
    recipe_id int references Recipes(id) on delete cascade not null
)

CREATE TABLE Ingredients
(
    id serail not null unique,
    title varchar(255) not null,
    description varchar(255) not null,
    meat boolean not null default false,
    seafood boolean not null default false
)

CREATE TABLE Recipe_Ingredients
(
    id serial not null unique,
    recipe_id int references Recipes(id) on delete cascade not null
    recipe_id int references Ingredients(id) on delete cascade not null,
)

