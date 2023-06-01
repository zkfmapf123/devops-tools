package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/*
go run main.go http://www.naver.com
*/

func main() {

	args := os.Args

	if len(args) < 2 {
		log.Fatalf("Usage : example) http://www.naver.com \n")
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		log.Fatalf("URL is in invalid format : %v", err)
	}

	res, err := http.Get(args[1])
	if err != nil {
		log.Fatalf("status : %v, reason : %v", res.Status, err)
	}

	defer res.Body.Close()

	// read until EOF
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("invalid error : %v", err)
	}

	fmt.Printf("body : %s", body)
}
