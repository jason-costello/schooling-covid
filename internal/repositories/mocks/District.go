// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/jason-costello/schooling-covid/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// District is an autogenerated mock type for the District type
type District struct {
	mock.Mock
}

// AllDistricts provides a mock function with given fields: ctx
func (_m *District) AllDistricts(ctx context.Context) ([]models.District, error) {
	ret := _m.Called(ctx)

	var r0 []models.District
	if rf, ok := ret.Get(0).(func(context.Context) []models.District); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.District)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DistrictByShortName provides a mock function with given fields: ctx, districtShortName
func (_m *District) DistrictByShortName(ctx context.Context, districtShortName string) (map[string]models.School, error) {
	ret := _m.Called(ctx, districtShortName)

	var r0 map[string]models.School
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]models.School); ok {
		r0 = rf(ctx, districtShortName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]models.School)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, districtShortName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetName provides a mock function with given fields: ctx, shortname
func (_m *District) GetName(ctx context.Context, shortname string) (string, error) {
	ret := _m.Called(ctx, shortname)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortname)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
