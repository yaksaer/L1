package main

import (
	"fmt"
	"reflect"
)

func main() {
	channel := make(chan string)
	typesList := []interface{}{1, "Hello", true, channel}
	for _, v := range typesList {
		fmt.Println(reflect.TypeOf(v))
	}
}
