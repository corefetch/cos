drop table if exists accounts;

create table accounts (
    "id" bigint,
    "login" varchar(64),
    "password" varchar(128),
    "names" varchar(128),
    "meta" json,
    "verified" TIMESTAMP,
    "created" TIMESTAMP WITH TIME ZONE DEFAULT now()
);