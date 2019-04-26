create table developers
(
    id            serial      not null
        constraint developers_pk
            primary key,
    name          varchar(20) not null,
    age           integer     not null,
    primary_skill varchar(20) not null
);
alter table developers
    owner to postgres;
create table customers
(
    id       serial      not null
        constraint customer_pk
            primary key,
    name     varchar(20) not null,
    money    integer     not null,
    discount integer     not null,
    state    varchar(10) not null
);

alter table customers
    owner to postgres;