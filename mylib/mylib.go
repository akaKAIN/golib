package mylib

import (
	"fmt"
	"time"
)

func ShowCurrentTime() {
	t := time.Now().Format("2006.01.02-15.04.05")
	fmt.Println(t)
}
