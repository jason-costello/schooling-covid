// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package storage

import (
	"context"
	"time"
)

const addCountData = `-- name: AddCountData :exec
Insert into counts(school_id, positive, symptomatic, exposed, created_at)
values($1, $2, $3, $4, now())
`

type AddCountDataParams struct {
	SchoolID    int32
	Positive    int32
	Symptomatic int32
	Exposed     int32
}

func (q *Queries) AddCountData(ctx context.Context, arg AddCountDataParams) error {
	_, err := q.db.ExecContext(ctx, addCountData,
		arg.SchoolID,
		arg.Positive,
		arg.Symptomatic,
		arg.Exposed,
	)
	return err
}

const addDistrictData = `-- name: AddDistrictData :exec
Insert into districts(name, short_name, created_at)
values($1, $2,  now())
`

type AddDistrictDataParams struct {
	Name      string
	ShortName string
}

func (q *Queries) AddDistrictData(ctx context.Context, arg AddDistrictDataParams) error {
	_, err := q.db.ExecContext(ctx, addDistrictData, arg.Name, arg.ShortName)
	return err
}

const addSchoolData = `-- name: AddSchoolData :exec
Insert into schools(name, short_name, district_id,created_at)
values($1, $2, $3, now())
`

type AddSchoolDataParams struct {
	Name       string
	ShortName  string
	DistrictID int32
}

func (q *Queries) AddSchoolData(ctx context.Context, arg AddSchoolDataParams) error {
	_, err := q.db.ExecContext(ctx, addSchoolData, arg.Name, arg.ShortName, arg.DistrictID)
	return err
}

const getAllCountsForSchoolID = `-- name: GetAllCountsForSchoolID :many
SELECT  s.name
     , c.positive as positive
     , c.symptomatic as synptomatic
     , c.exposed as exposed
     , date(c.created_at) as observation_date
from schools as s
         left join counts as c on  s.id = c.school_id
where s.id = $1
`

type GetAllCountsForSchoolIDRow struct {
	Name            string
	Positive        int32
	Synptomatic     int32
	Exposed         int32
	ObservationDate time.Time
}

func (q *Queries) GetAllCountsForSchoolID(ctx context.Context, id int32) ([]GetAllCountsForSchoolIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllCountsForSchoolID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCountsForSchoolIDRow
	for rows.Next() {
		var i GetAllCountsForSchoolIDRow
		if err := rows.Scan(
			&i.Name,
			&i.Positive,
			&i.Synptomatic,
			&i.Exposed,
			&i.ObservationDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllDistricts = `-- name: GetAllDistricts :many

SELECT id, name, short_name, created_at, updated_at from districts order by name
`

// query.sql
func (q *Queries) GetAllDistricts(ctx context.Context) ([]District, error) {
	rows, err := q.db.QueryContext(ctx, getAllDistricts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []District
	for rows.Next() {
		var i District
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ShortName,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSchools = `-- name: GetAllSchools :many
SELECT id, name, short_name, district_id, created_at, updated_at from schools
`

func (q *Queries) GetAllSchools(ctx context.Context) ([]School, error) {
	rows, err := q.db.QueryContext(ctx, getAllSchools)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []School
	for rows.Next() {
		var i School
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ShortName,
			&i.DistrictID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllSchoolsForDistrict = `-- name: GetAllSchoolsForDistrict :many
SELECT id, name, short_name, district_id, created_at, updated_at from schools where district_id = $1
`

func (q *Queries) GetAllSchoolsForDistrict(ctx context.Context, districtID int32) ([]School, error) {
	rows, err := q.db.QueryContext(ctx, getAllSchoolsForDistrict, districtID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []School
	for rows.Next() {
		var i School
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ShortName,
			&i.DistrictID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountData = `-- name: GetCountData :one
select id, school_id, positive, symptomatic, exposed, created_at, updated_at from counts where id = $1
`

func (q *Queries) GetCountData(ctx context.Context, id int32) (Count, error) {
	row := q.db.QueryRowContext(ctx, getCountData, id)
	var i Count
	err := row.Scan(
		&i.ID,
		&i.SchoolID,
		&i.Positive,
		&i.Symptomatic,
		&i.Exposed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCountsBySchoolIDByDate = `-- name: GetCountsBySchoolIDByDate :many
SELECT  s.name
     ,c.positive as positive
     ,c.symptomatic as synptomatic
     ,c.exposed as exposed
     , date(c.created_at) as observation_date
from schools as s
         left join counts as c on  s.id = c.school_id
where s.id = $1 and date(c.created_at) = $2
`

type GetCountsBySchoolIDByDateParams struct {
	ID        int32
	CreatedAt time.Time
}

type GetCountsBySchoolIDByDateRow struct {
	Name            string
	Positive        int32
	Synptomatic     int32
	Exposed         int32
	ObservationDate time.Time
}

func (q *Queries) GetCountsBySchoolIDByDate(ctx context.Context, arg GetCountsBySchoolIDByDateParams) ([]GetCountsBySchoolIDByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getCountsBySchoolIDByDate, arg.ID, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCountsBySchoolIDByDateRow
	for rows.Next() {
		var i GetCountsBySchoolIDByDateRow
		if err := rows.Scan(
			&i.Name,
			&i.Positive,
			&i.Synptomatic,
			&i.Exposed,
			&i.ObservationDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDailyCountSummaryByDistrictID = `-- name: GetDailyCountSummaryByDistrictID :one
SELECT  d.name
     , sum(c.positive) as positive
     , sum(c.symptomatic) as synptomatic
     , sum(c.exposed) as exposed
     , date(c.created_at) as observation_date
from districts as d
         left join schools as s on  d.id = s.district_id
         left join counts as c on  s.id = c.school_id
where d.id = $1
group by d.name, date(c.created_at)
`

type GetDailyCountSummaryByDistrictIDRow struct {
	Name            string
	Positive        int64
	Synptomatic     int64
	Exposed         int64
	ObservationDate time.Time
}

func (q *Queries) GetDailyCountSummaryByDistrictID(ctx context.Context, id int32) (GetDailyCountSummaryByDistrictIDRow, error) {
	row := q.db.QueryRowContext(ctx, getDailyCountSummaryByDistrictID, id)
	var i GetDailyCountSummaryByDistrictIDRow
	err := row.Scan(
		&i.Name,
		&i.Positive,
		&i.Synptomatic,
		&i.Exposed,
		&i.ObservationDate,
	)
	return i, err
}

const getDailyCountSummaryPerSchoolByDistrictID = `-- name: GetDailyCountSummaryPerSchoolByDistrictID :many
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
group by d.name, s.name, date(c.created_at)
`

type GetDailyCountSummaryPerSchoolByDistrictIDRow struct {
	Name            string
	Name_2          string
	Positive        int64
	Synptomatic     int64
	Exposed         int64
	ObservationDate time.Time
}

func (q *Queries) GetDailyCountSummaryPerSchoolByDistrictID(ctx context.Context) ([]GetDailyCountSummaryPerSchoolByDistrictIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getDailyCountSummaryPerSchoolByDistrictID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetDailyCountSummaryPerSchoolByDistrictIDRow
	for rows.Next() {
		var i GetDailyCountSummaryPerSchoolByDistrictIDRow
		if err := rows.Scan(
			&i.Name,
			&i.Name_2,
			&i.Positive,
			&i.Synptomatic,
			&i.Exposed,
			&i.ObservationDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDistrict = `-- name: GetDistrict :one
SELECT id, name, short_name, created_at, updated_at from districts where id = $1
`

func (q *Queries) GetDistrict(ctx context.Context, id int32) (District, error) {
	row := q.db.QueryRowContext(ctx, getDistrict, id)
	var i District
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ShortName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getDistrictData = `-- name: GetDistrictData :one
select d.id, d.name, d.short_name
from districts as d

where d.id = $1
`

type GetDistrictDataRow struct {
	ID        int32
	Name      string
	ShortName string
}

func (q *Queries) GetDistrictData(ctx context.Context, id int32) (GetDistrictDataRow, error) {
	row := q.db.QueryRowContext(ctx, getDistrictData, id)
	var i GetDistrictDataRow
	err := row.Scan(&i.ID, &i.Name, &i.ShortName)
	return i, err
}

const getSchoolByID = `-- name: GetSchoolByID :one
SELECT id, name, short_name, district_id, created_at, updated_at from schools where id = $1
`

func (q *Queries) GetSchoolByID(ctx context.Context, id int32) (School, error) {
	row := q.db.QueryRowContext(ctx, getSchoolByID, id)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ShortName,
		&i.DistrictID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSchoolData = `-- name: GetSchoolData :one
select id, name, short_name, district_id, created_at, updated_at from schools where id = $1
`

func (q *Queries) GetSchoolData(ctx context.Context, id int32) (School, error) {
	row := q.db.QueryRowContext(ctx, getSchoolData, id)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ShortName,
		&i.DistrictID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateCountData = `-- name: UpdateCountData :exec
Update counts set school_id = $1, positive = $2, exposed = $3, symptomatic = $4
where id = $5
`

type UpdateCountDataParams struct {
	SchoolID    int32
	Positive    int32
	Exposed     int32
	Symptomatic int32
	ID          int32
}

func (q *Queries) UpdateCountData(ctx context.Context, arg UpdateCountDataParams) error {
	_, err := q.db.ExecContext(ctx, updateCountData,
		arg.SchoolID,
		arg.Positive,
		arg.Exposed,
		arg.Symptomatic,
		arg.ID,
	)
	return err
}

const updateDistrictData = `-- name: UpdateDistrictData :exec
Update districts set name = $1, short_name = $2
where id = $3
`

type UpdateDistrictDataParams struct {
	Name      string
	ShortName string
	ID        int32
}

func (q *Queries) UpdateDistrictData(ctx context.Context, arg UpdateDistrictDataParams) error {
	_, err := q.db.ExecContext(ctx, updateDistrictData, arg.Name, arg.ShortName, arg.ID)
	return err
}

const updateSchoolData = `-- name: UpdateSchoolData :exec
Update schools set name = $1, short_name = $2, district_id = $3
where id = $4
`

type UpdateSchoolDataParams struct {
	Name       string
	ShortName  string
	DistrictID int32
	ID         int32
}

func (q *Queries) UpdateSchoolData(ctx context.Context, arg UpdateSchoolDataParams) error {
	_, err := q.db.ExecContext(ctx, updateSchoolData,
		arg.Name,
		arg.ShortName,
		arg.DistrictID,
		arg.ID,
	)
	return err
}
