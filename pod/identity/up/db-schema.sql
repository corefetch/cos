drop table if exists accounts;

create table accounts (
    "id" bigint,
    "login" varchar(64),
    "password" varchar(128),
    "names" varchar(128),
    "verified" TIMESTAMP,
    "created" TIMESTAMP WITH TIME ZONE DEFAULT now(),
    UNIQUE("login"),
    PRIMARY KEY("id")
);

drop table if exists meta;

create table meta (
    "account" bigint,
    "name" varchar(32),
    "value" varchar(128),
    UNIQUE("account", "name")
);