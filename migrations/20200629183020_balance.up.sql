CREATE TABLE users (
    id varchar(36) NOT NULL UNIQUE,
    email varchar(50) NOT NULL UNIQUE,
    "password" varchar(60) NOT NULL,
    username varchar(50),
    admin boolean NOT NULL DEFAULT false,
    status boolean NOT NULL DEFAULT true,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text,
    CONSTRAINT pk_users PRIMARY KEY (id)
    );
-----------------------------------------------------------------------------------------------------------
CREATE TABLE stocks (
    id serial,
    owner varchar(36) REFERENCES users (id) ON DELETE CASCADE,
    name varchar(100) NOT NULL,
    status boolean NOT NULL DEFAULT true,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text,
    CONSTRAINT pk_stocks PRIMARY KEY (id)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE client_type (
    type integer,
    short_name varchar(5) NOT NULL,
    description varchar(50) NOT NULL,
    CONSTRAINT pk_client_type PRIMARY KEY (type)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE clients (
    id serial,
    name varchar(200) NOT NULL,
    "user" varchar(36) REFERENCES users (id) ON DELETE CASCADE,
    type integer REFERENCES client_type (type) ON DELETE RESTRICT,
    markup real,
    status boolean NOT NULL DEFAULT true,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text, --TODO: можно добавить поле реквизиты и ссылаться на таблицу с данными:телефон-адрес и т.д.
    CONSTRAINT pk_clients PRIMARY KEY (id)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE document_type (
    type integer,
    short_name varchar(5) NOT NULL,
    description varchar(50) NOT NULL,
    CONSTRAINT pk_document_type PRIMARY KEY (type)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE documents (
    id bigserial,
    "user" varchar(36) REFERENCES users (id) ON DELETE CASCADE,
    client  integer REFERENCES clients (id) ON DELETE CASCADE,
    type integer REFERENCES document_type (type) ON DELETE RESTRICT,
    status boolean NOT NULL DEFAULT true,
    summa NUMERIC(8,2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    comment text,
    path_to_file text,
    CONSTRAINT pk_documents PRIMARY KEY (id)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE products (
    id bigserial,
    number varchar(50) NOT NULL,
    producer varchar(50) NOT NULL,
    description varchar(100),
    CONSTRAINT pk_products PRIMARY KEY (id)
);
-----------------------------------------------------------------------------------------------------------
CREATE TABLE products_in_stocks(
    stock integer REFERENCES stocks (id) ON DELETE RESTRICT,
    product bigint REFERENCES products (id) ON DELETE RESTRICT,
    quantity integer,
    price NUMERIC(8,2),
    UNIQUE (stock, product, price) --ON CONFLICT UNIQUE DO UPDATE SET quantity++
);

-----------------------------------------------------------------------------------------------------------
CREATE TABLE products_in_documents(
    document bigint REFERENCES documents (id) ON DELETE CASCADE,
    product bigint REFERENCES products (id) ON DELETE RESTRICT,
    quantity integer,
    price NUMERIC(8,2)
);
-----------------------------------------------------------------------------------------------------------
