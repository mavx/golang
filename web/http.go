package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

func main() {
	res, _ := http.Get("https://api.mybitx.com/api/1/ticker?pair=XBTMYR")
	fmt.Println(res)
	fmt.Println(res.Body)
	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
	fmt.Println(res.Header)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	var s string
	test := json.Unmarshal(body, &s)

	fmt.Println(test)
}
