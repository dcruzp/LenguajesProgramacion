package main

import (
	"fmt"
)

type T interface{}

func PrintAny(any T) {
	switch any.(type) {
	case string:
		println(any)

	case int:
		fmt.Printf("%d", any)
	}
}
