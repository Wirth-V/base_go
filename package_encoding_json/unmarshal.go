package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

	type myStruct struct {
		Name   string
		Age    int
		Status bool
		Values []int
	}

	var s myStruct

	if err := json.Unmarshal(data, &s); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
	fmt.Printf("\n")
	fmt.Printf("%s\n", s)

}
