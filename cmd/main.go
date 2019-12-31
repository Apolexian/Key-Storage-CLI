package main

import (
	"../vault"
	"fmt"
)

func main() {
	v := vault.Memory("my-test-key")
	err := v.Set("demo", "test_value")

	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo")

	if err != nil {
		panic(err)
	}

	fmt.Println("Plain value:", plain)
}
