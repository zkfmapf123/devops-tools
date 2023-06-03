package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"zkfmapf123/parse-json/utils"
)

var DATA_FILE = "data.json"

type Datas struct {
	Name           string   `json:"name"`
	JobTitle       string   `json:"job_title"`
	Skills         []string `json:"skills"`
	JobPersonliaty string   `json:"job_personality"`
}

func main() {
	b, err := os.ReadFile(DATA_FILE)

	if err != nil {
		log.Fatalln(err)
	}

	var data Datas
	err = json.Unmarshal(b, &data)

	if err != nil {
		log.Fatalln(err)
	}

	b, _ = json.Marshal(data)
	fmt.Printf("%s\n", b)

	// Object Method
	object := utils.JsonObject{}
	fmt.Println("keys >> ", object.Keys(data))
	fmt.Println("values >> ", object.Values(data))
	fmt.Println("entires >> ", object.Entries(data))
}
