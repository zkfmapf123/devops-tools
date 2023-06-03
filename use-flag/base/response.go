package base

import (
	"fmt"
	"strings"
)

type Response interface {
	GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

func (w Words) GetResponse() string {

	return fmt.Sprintf("Words: %s >>  ", strings.Join(w.Words, ", "))
}

func (o Occurrence) GetResponse() string {
	words := []string{}

	for w, o := range o.Words {
		words = append(words, fmt.Sprintf("%s (%d)", w, o))
	}

	return fmt.Sprintf("Words: %s", strings.Join(words, ", "))
}
