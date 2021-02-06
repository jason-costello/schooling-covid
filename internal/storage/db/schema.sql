create table schools
(
    id serial not null
        constraint schools_pk
        primary key,
    name varchar(500) not null,
    short_name varchar(50) not null,
    district_id integer not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);

alter table schools owner to gzlgiwbcviwknc;

create unique index schools_id_uindex
    on schools (id);

create unique index schools_short_name_uindex
    on schools (short_name);

create table districts
(
    id serial not null
        constraint districts_pk
        primary key,
    name varchar(500) not null,
    short_name varchar(50) not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);

alter table districts owner to gzlgiwbcviwknc;

create table counts
(
    id serial not null
        constraint counts_pk
        primary key,
    school_id integer not null,
    positive integer default 0 not null,
    symptomatic integer default 0 not null,
    exposed integer default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);

alter table counts owner to gzlgiwbcviwknc;

create function trigger_set_timestamp() returns trigger
    language plpgsql
as $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$;

alter function trigger_set_timestamp() owner to gzlgiwbcviwknc;

create trigger set_timestamp
    before update
    on schools
    for each row
    execute procedure trigger_set_timestamp();

create trigger set_timestamp
    before update
    on districts
    for each row
    execute procedure trigger_set_timestamp();

create trigger set_timestamp
    before update
    on counts
    for each row
    execute procedure trigger_set_timestamp();

