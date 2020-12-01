package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	response, err := http.Get("https://httpbin.org/get")
	exitOnErr(err, "http Get request failed")

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	exitOnErr(err, "Failed to read response body")

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
	fmt.Printf("StatusCode: %d\n", response.StatusCode)
	fmt.Printf("Headers: %s\n", response.Header)

	fmt.Printf("Header: Content Type - %s", response.Header["Content-Type"])
}

func exitOnErr(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
