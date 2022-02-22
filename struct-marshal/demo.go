package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name",omitempty`
}

func main() {
	a := &Person{
		Age:  10,
		Name: "win",
	}

	b := &struct {
		Age  int     `json:"age,omitempty"`
		Name *string `json:"name,omitempty"`
	}{
		Age: a.Age,
	}

	if len(a.Name) > 0 {
		b.Name = &a.Name
	}

	data, err := json.Marshal(b)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
