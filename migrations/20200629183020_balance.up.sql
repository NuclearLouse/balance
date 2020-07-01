CREATE TABLE users (
    id serial,
    email varchar(50) NOT NULL UNIQUE,
    "password" varchar(60) NOT NULL,
    is_admin boolean NOT NULL DEFAULT false,
    status boolean NOT NULL DEFAULT true,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text
    );
-----------------------------------------------------------------------------------------------------------