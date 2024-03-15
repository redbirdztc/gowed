package gowed

import (
	"reflect"
	"testing"

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

func TestExtracMapUnexpiredData(t *testing.T) {
	type ExpirableStruct struct {
		Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	m := map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}
	result := ExtracMapUnexpiredData(m)
	expected := []ExpirableStruct{unexpirable}

	assert.True(t, reflect.DeepEqual(result, expected), "result: %v, expected: %v", result, expected)
}

func TestExtracMapExpiredData(t *testing.T) {
	type ExpirableStruct struct {
		Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	m := map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}
	result := ExtracMapExpiredData(m)
	expected := []ExpirableStruct{expirable}

	assert.True(t, reflect.DeepEqual(result, expected), "result: %v, expected: %v", result, expected)
}

func TestRemoveExpiredData(t *testing.T) {
	type ExpirableStruct struct {
		Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	m := map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}
	RemoveExpiredData(m)
	expected := map[int]ExpirableStruct{
		1: unexpirable,
	}

	assert.True(t, reflect.DeepEqual(m, expected), "result: %v, expected: %v", m, expected)
}

func TestRemoveUnexpiredData(t *testing.T) {
	type ExpirableStruct struct {
		Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	m := map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}
	RemoveUnexpiredData(m)
	expected := map[int]ExpirableStruct{
		2: expirable,
	}

	assert.True(t, reflect.DeepEqual(m, expected), "result: %v, expected: %v", m, expected)
}
