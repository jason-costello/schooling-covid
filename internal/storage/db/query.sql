-- query.sql

-- name: GetAllDistricts :many
SELECT id, name, short_name, created_at, updated_at from districts order by name;

-- name: GetDistrict :one
SELECT id, name, short_name, created_at, updated_at from districts where id = $1;

-- name: GetAllSchools :many
SELECT id, name, short_name, district_id, created_at, updated_at from schools;

-- name: GetAllSchoolsForDistrict :many
SELECT id, name, short_name, district_id, created_at, updated_at from schools where district_id = $1;

-- name: GetSchoolByID :one
SELECT id, name, short_name, district_id, created_at, updated_at from schools where id = $1;


-- name: GetDailyCountSummaryByDistrictID :one
SELECT  d.name
     , sum(c.positive) as positive
     , sum(c.symptomatic) as synptomatic
     , sum(c.exposed) as exposed
     , date(c.created_at) as observation_date
from districts as d
         left join schools as s on  d.id = s.district_id
         left join counts as c on  s.id = c.school_id
where d.id = $1
group by d.name, date(c.created_at);


-- name: GetDailyCountSummaryPerSchoolByDistrictID :many
SELECT  d.name
     ,s.name
     , sum(c.positive) as positive
     , sum(c.symptomatic) as synptomatic
     , sum(c.exposed) as exposed
     , date(c.created_at) as observation_date
from districts as d
         left join schools as s on  d.id = s.district_id
         left join counts as c on  s.id = c.school_id
where d.id = 1
group by d.name, s.name, date(c.created_at);


-- name: GetAllCountsForSchoolID :many
SELECT  s.name
     , c.positive as positive
     , c.symptomatic as synptomatic
     , c.exposed as exposed
     , date(c.created_at) as observation_date
from schools as s
         left join counts as c on  s.id = c.school_id
where s.id = $1;


-- name: GetCountsBySchoolIDByDate :many
SELECT  s.name
     ,c.positive as positive
     ,c.symptomatic as synptomatic
     ,c.exposed as exposed
     , date(c.created_at) as observation_date
from schools as s
         left join counts as c on  s.id = c.school_id
where s.id = $1 and date(c.created_at) = $2;




-- name: AddCountData :exec
Insert into counts(school_id, positive, symptomatic, exposed, created_at)
values($1, $2, $3, $4, now());

-- name: UpdateCountData :exec
Update counts set school_id = $1, positive = $2, exposed = $3, symptomatic = $4
where id = $5;


-- name: GetCountData :one
select * from counts where id = $1;

-- name: AddDistrictData :exec
Insert into districts(name, short_name, created_at)
values($1, $2,  now());

-- name: UpdateDistrictData :exec
Update districts set name = $1, short_name = $2
where id = $3;

-- name: GetDistrictData :one
select d.id, d.name, d.short_name
from districts as d

where d.id = $1;



-- name: AddSchoolData :exec
Insert into schools(name, short_name, district_id,created_at)
values($1, $2, $3, now());

-- name: UpdateSchoolData :exec
Update schools set name = $1, short_name = $2, district_id = $3
where id = $4;

-- name: GetSchoolData :one
select * from schools where id = $1;
