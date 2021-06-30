\c golinks

SET TIMEZONE='Brazil/East';

CREATE TABLE "golinks"(
    "original_url" varchar(255) primary key,
    "golink" varchar(255) not null,
    "creation_date" timestamp without time zone default current_timestamp
);
