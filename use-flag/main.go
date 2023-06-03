package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"zkfmapf123/use-flag/base"
	"zkfmapf123/use-flag/utils"
)

var (
	requestURL string
	password   string
	parseURL   *url.URL
	err        error
)

func main() {

	flag.StringVar(&requestURL, "url", "", "url to access")
	flag.StringVar(&password, "password", "", "use a password to accecss our api")
	flag.Parse()

	if parseURL, err = url.ParseRequestURI(requestURL); err != nil {
		log.Fatalf("Help: ./http-get -h/nURL is not valid URL %s\n", requestURL)
	}

	res, err := doRequest(parseURL.String())
	fmt.Println(res, err)
}

func doRequest(url string) (base.Response, error) {

	res, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("Get Error : %s", err)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll Error : %s\n", err)
	}

	if res.Status != "200" {
		return nil, fmt.Errorf("Invalid output %s\n", err)
	}

	if !json.Valid(b) {
		return nil, utils.ResponseError{
			StatusCode: res.Status,
			Body:       string(b),
			Err:        "Invalid Json Object",
		}
	}

	// Other job

	return nil, nil
}
