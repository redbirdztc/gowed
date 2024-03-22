package main

import (
	"fmt"
	"reflect"

	"github.com/redbirdztc/gowed"
)

type Expired struct{}

func (Expired) IsExpired() bool {
	return true
}

type Unexpired struct{}

func (Unexpired) IsExpired() bool {
	return false
}

func main() {
	type ExpirableStruct struct {
		gowed.Expirer
	}

	expirable := ExpirableStruct{
		Expirer: Expired{},
	}
	unexpirable := ExpirableStruct{
		Expirer: Unexpired{},
	}

	s := []ExpirableStruct{
		unexpirable,
		expirable,
	}

	result := gowed.ExtracSliceUnexpiredData(s)
	fmt.Println("result of ExtracMapUnexpiredData: ")
	for _, v := range result {
		fmt.Println("\t", reflect.TypeOf(v.Expirer))
	}

	result = gowed.ExtracSliceExpiredData(s)
	fmt.Println("result of ExtracMapExpiredData: ")
	for _, v := range result {
		fmt.Println("\t", reflect.TypeOf(v.Expirer))
	}
}
