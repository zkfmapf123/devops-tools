package main

import (
	"fmt"
	"zkfmapf123/custom-error/error"
)

func main() {

	err := error.RequestError{
		HTTPCode: 200,
		Body:     "hello world",
		Err:      "not error",
	}

	fmt.Println(err)
}
