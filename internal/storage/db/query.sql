-- query.sql

-- name: GetAllDistricts :many
SELECT  name, short_name, short_name, updated_at from districts order by name;

-- name: GetDistrict :one
SELECT name, short_name, short_name, updated_at from districts where short_name = $1;

-- name: GetAllSchools :many
SELECT  name, short_name, district_short_name, created_at, updated_at from schools;

-- name: GetAllSchoolsForDistrict :many
SELECT  d.name, d.short_name, s.name
     , s.short_name, s.district_short_name
     , s.created_at, s.updated_at
    , c.positive, c.symptomatic, c.exposed, c.count_date

from districts as d
    join schools as s
on d.short_name = s.district_short_name
join counts as c
on s.short_name = c.school_short_name
where district_short_name = $1
Order by d.short_name, s.short_name, c.count_date;

-- name: GetSchoolByID :one
SELECT  name, short_name, district_short_name, created_at, updated_at from schools where short_name = $1;

