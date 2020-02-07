package main

import "fmt"

func main() {
	m := map[string]int{"foo": 42, "boo": 27}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 17
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}