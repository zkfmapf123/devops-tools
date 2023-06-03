package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type MySlowReader struct {
	FilePath string
}

func (m MySlowReader) Read(p []byte) (n int, err error) {

	file, _ := os.Open(m.FilePath)

	file.Read(p)
	fmt.Println(p)

	return 0, io.EOF
}

func main() {

	filePath := "data.txt"

	benchmarkFileRead("benchmakr test", func() {
		for i := 0; i < 100000; i++ {

		}
	})

	benchmarkFileRead("Normal Read File", func() {
		_, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}
	})

	benchmarkFileRead("use better performance read file", func() {
		file, _ := os.Open(filePath)

		sc := bufio.NewScanner(file)
		for sc.Scan() {
			// read
		}

		defer file.Close()
	})

	msr := MySlowReader{
		FilePath: filePath,
	}

	b, err := ioutil.ReadAll(msr)
	fmt.Println(b, err)

}

func benchmarkFileRead(text string, f func()) {
	start := time.Now()
	f()
	fmt.Println(text, " >> ", time.Since(start))
}
