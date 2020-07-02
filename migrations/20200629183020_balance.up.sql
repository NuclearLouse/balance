CREATE TABLE users (
    id varchar(36) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    "password" varchar(60) NOT NULL,
    username varchar(50),
    is_admin boolean NOT NULL DEFAULT false,
    status boolean NOT NULL DEFAULT true,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text,
    CONSTRAINT pk_users PRIMARY KEY (id)
    );
-----------------------------------------------------------------------------------------------------------