package main

import (
	"fmt"
	"reflect"
)

type Restaurant struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	SearchName string `json:searchName`
	Type       string `json:type`
}

func main() {
	fmt.Println("Hello, playground")
	res := Restaurant{}
	resReflect := reflect.ValueOf(&res).Elem()
	resReflectType := resReflect.Type()

	for i := 0; i < resReflect.NumField(); i++ {
		t := resReflectType.Field(i).Name
		fmt.Println("field: ", t)
	}

}
