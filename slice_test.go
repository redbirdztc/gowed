package gowed_test

import (
	"testing"

	"github.com/redbirdztc/gowed"
	"github.com/stretchr/testify/assert"
)

func TestExtracSliceUnexpiredData(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	slice := []ExpirableStruct{unexpirable, expirable}
	result := gowed.ExtracSliceUnexpiredData(slice)
	assert.Len(t, result, 1)
	assert.Equal(t, unexpirable, result[0])
}

func TestExtracSliceExpiredData(t *testing.T) {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	slice := []ExpirableStruct{unexpirable, expirable}
	result := gowed.ExtracSliceExpiredData(slice)
	assert.Len(t, result, 1)
	assert.Equal(t, expirable, result[0])
}
