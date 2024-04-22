package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	s := myStruct{
		Name:   "John Connor",
		Age:    35,
		Status: true,
		Values: []int{15, 11, 37},
	}

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
	fmt.Printf("\n")
	fmt.Printf("%s", data)

}
