create table developers
(
  id           serial      not null
    constraint developers_pk
      primary key,
  name         varchar(20) not null,
  age          integer     not null,
  primary_skill varchar(20) not null
);

alter table developers
  owner to postgres;

