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

	m := map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}

	result := gowed.ExtracMapUnexpiredData(m)
	fmt.Println("result of ExtracMapUnexpiredData: ")
	for _, v := range result {
		fmt.Println("\t", reflect.TypeOf(v.Expirer))
	}

	result = gowed.ExtracMapExpiredData(m)
	fmt.Println("result of ExtracMapExpiredData: ")
	for _, v := range result {
		fmt.Println("\t", reflect.TypeOf(v.Expirer))
	}

	gowed.RemoveMapExpiredData(m)
	fmt.Println("result of RemoveExpiredData: ")
	for k, v := range m {
		fmt.Println("\t", "key:", k, " valueType:", reflect.TypeOf(v.Expirer))
	}

	m = map[int]ExpirableStruct{
		1: unexpirable,
		2: expirable,
	}
	gowed.RemoveMapUnexpiredData(m)
	fmt.Println("result of RemoveUnexpiredData: ")
	for k, v := range m {
		fmt.Println("\t", "key:", k, " valueType:", reflect.TypeOf(v.Expirer))
	}
}
