package repositories

import (
	"context"
	"errors"
	"github.com/jason-costello/schooling-covid/internal/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


func TestDistrictRepository_GetName(t *testing.T) {
	md := mocks.District{}

	md.On("GetName", mock.MatchedBy(func(ctx context.Context) bool { return true}), mock.AnythingOfType("string")).Return(
		func(ctx context.Context, s string) string {
			switch s {
			case "NBISD":
				return "New Braunfels ISD"
			case "CISD":
				return "Comal ISD"
			default:
				return ""
			}
		},
		func(ctx context.Context, s string) error {
			switch s{
			case "NBISD":
				return  nil
			case "CISD":
				return nil
			default:
				return errors.New("not found")
			}
		})


// var ctx context.Context
var ctx context.Context
longName, err := md.GetName(ctx, "NBISD")
assert.NoError(t, err)
assert.Equal(t, "New Braunfels ISD", longName)


longName = ""
err = nil
	longName, err = md.GetName(ctx, "CISD")
	assert.NoError(t, err)
	assert.Equal(t, "Comal ISD", longName)

	longName = ""
	err = nil
	longName, err = md.GetName(ctx, "")
	assert.Error(t, err)
	assert.Equal(t, "", longName)

	longName = ""
	err = nil
	longName, err = md.GetName(ctx, "not a real name and should have a blank return with an error")
	assert.Error(t, err)
	assert.Equal(t, "", longName)

}

