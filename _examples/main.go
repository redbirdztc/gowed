package main

import (
	"fmt"
	"time"

	"github.com/redbirdztc/gowed"
)

// Implement `Expirer`
type Expirer struct {
	ExpirerAt *time.Time
}

var _ gowed.Expirer = (*Expirer)(nil)

func (e *Expirer) IsExpired() bool {
	if e.ExpirerAt == nil {
		return false
	}
	now := time.Now()
	return now.After(*e.ExpirerAt)
}

func main() {
	t := time.Now().AddDate(0, 0, -1)
	obj := &Expirer{ExpirerAt: &t}

	// `HasExpiredData` will check if the data has expired
	if gowed.HasExpiredData(obj) {
		fmt.Println("obj has expired with expired time: ", t)
	} else {
		fmt.Println("obj has not expired with expired time: ", t)
	}

	t=time.Now().AddDate(0, 0, 1)
	obj.ExpirerAt = &t

	// `HasExpiredData` will check if the data has expired
	if gowed.HasExpiredData(obj) {
		fmt.Println("obj has expired with expired time: ", t)
	} else {
		fmt.Println("obj has not expired with expired time: ", t)
	}

	// As to basic types, it will return false
	if gowed.HasExpiredData(1) {
		fmt.Println("1 has expired")
	} else {
		fmt.Println("1 has not expired")
	}

	// As to nil, it will return false
	if gowed.HasExpiredData(nil) {
		fmt.Println("nil has expired")
	} else {
		fmt.Println("nil has not expired")
	}

}
