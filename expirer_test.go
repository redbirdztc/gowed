package gowed_test

import (
	"testing"

	"github.com/redbirdztc/gowed"
	"github.com/stretchr/testify/assert"
)

type Expired struct{}

func (Expired) IsExpired() bool {
	return true
}

type Unexpired struct{}

func (Unexpired) IsExpired() bool {
	return false
}

func TestHasExpiredData_Nil(t *testing.T) {
	assert.False(t, gowed.HasExpiredData(nil))
	assert.False(t, gowed.HasExpiredData(any(nil)))
}
func TestHasExpiredData_Expirable(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	assert.True(t, gowed.HasExpiredData(expirable))
	assert.False(t, gowed.HasExpiredData(unexpirable))

}

func TestHasExpiredData_Array(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	slice := [4]any{unexpirable, 0, "2", byte(1)}
	assert.False(t, gowed.HasExpiredData(slice))

	slice[3] = expirable
	assert.True(t, gowed.HasExpiredData(slice))
}

func TestHasExpiredData_Slice(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	slice := []any{unexpirable, 0, "2", byte(1), 0.0}
	assert.False(t, gowed.HasExpiredData(slice))

	slice = append(slice, expirable)
	assert.True(t, gowed.HasExpiredData(slice))
}

func TestHasExpiredData_Map(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	m := map[int]any{}
	m[1] = 1
	m[2] = unexpirable
	assert.False(t, gowed.HasExpiredData(m))

	m[3] = expirable
	assert.True(t, gowed.HasExpiredData(m))
}
