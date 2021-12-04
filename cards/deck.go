package main

import (
	"database/sql/driver"
	"fmt"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}