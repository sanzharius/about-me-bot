package main

import (
	"fmt"
	"time"
)

func main() {
	ConvertTime()
}

func ConvertTime() string {
	l1 := "03:04:05PM"
	l2 := "15:04:05"
	t, err := time.Parse(l1, "07:05:45PM")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(t.Format(l2))
	return t.Format(l2)

}
