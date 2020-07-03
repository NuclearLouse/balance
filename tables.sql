-----------------------------------------------------------------------------------------------------------
CREATE TABLE "avail_stocks" (
    "name_stock" varchar(25),
    "part_num" varchar(25),
    "brand" varchar(25),
    "description" varchar(50),
    "count" float8,
    "price" float8,
    "comment" varchar(255) , 
    "markup" real, 
    "neg_margin" boolean
    );
-----------------------------------------------------------------------------------------------------------
CREATE TABLE "clients" (
    "id" integer primary key autoincrement,
    "name" varchar(30),
    "type" varchar(10),
    "user_owner" varchar(36),
    "status" boolean,
    "markup" float8,
    "created_at" datetime,
    "comment" varchar(255)
    );
-----------------------------------------------------------------------------------------------------------
CREATE TABLE "documents" (
    "id" integer primary key autoincrement,
    "number_doc" varchar(25),
    "user_owner" varchar(36),
    "summ" float8,
    "status" boolean,
    "type_doc" varchar(5),
    "client" varchar(30),
    "state" varchar(25),
    "file_xls" blob,
    "created_at" datetime,
    "comment" varchar(255)
    );
-----------------------------------------------------------------------------------------------------------
CREATE TABLE "invoices" (
    "number_source" varchar(25),
    "date_source" varchar(25),
    "number_doc" varchar(25),
    "part_num" varchar(25),
    "brand" varchar(25),
    "description" varchar(50),
    "count" float8,
    "price" float8,
    "stock" varchar(25)
    );
-----------------------------------------------------------------------------------------------------------