package main

import "fmt"

func main() {

	fmt.Println("Hello world")
}

/*
   go tool dist list

   // Linux -> Windows
   // Mac -> Windows
   GOOS=windows GOARCH=amd64 go build -o myprogram.exe main.go

   // Windows -> Linux
   SET GOOS=linux
   SET GOARCH=amd64
   go build -o myprogram main.go
*/
