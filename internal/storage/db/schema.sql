
create table districts
(
    name varchar(500) not null,
    short_name varchar(50) not null
        constraint districts_pk
        primary key,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null
);

alter table districts owner to gzlgiwbcviwknc;

create table schools
(
    name varchar(500) not null,
    short_name varchar(50) not null
        constraint schools_pk
        primary key,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null,
    district_short_name varchar(50) not null
        constraint schools_districts_short_name_fk
        references districts
);

alter table schools owner to gzlgiwbcviwknc;

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

create table counts
(
    positive integer default 0 not null,
    symptomatic integer default 0 not null,
    exposed integer default 0 not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updated_at timestamp default CURRENT_TIMESTAMP not null,
    school_short_name varchar(50) default 'none'::character varying not null
        constraint counts_schools_short_name_fk
        references schools,
    count_date varchar(10) not null,
    constraint date_school_pk
        primary key (school_short_name, count_date)
);

alter table counts owner to gzlgiwbcviwknc;

create trigger set_timestamp
    before update
    on counts
    for each row
    execute procedure trigger_set_timestamp();

create table collection_jobs
(
    id serial not null,
    district_short_name varchar(10) not null,
    school_short_name text not null,
    url text not null,
    created_at timestamp default CURRENT_TIMESTAMP not null,
    updtaed_at timestamp default CURRENT_TIMESTAMP not null,
    constraint district_school
        primary key (district_short_name, school_short_name)
);

alter table collection_jobs owner to gzlgiwbcviwknc;

create trigger set_timestamp
    before update
    on collection_jobs
    for each row
    execute procedure trigger_set_timestamp();


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

create trigger set_timestamp
    before update
    on collection_jobs
    for each row
    execute procedure trigger_set_timestamp();



select school_short_name, count_date, positive, symptomatic, exposed from counts order by school_short_name, count_date;
