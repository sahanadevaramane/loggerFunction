package main

import (
	"fmt"
	"github.com/sahanadevaramane/datadologger/helper"
	"io/ioutil"
	"net/http"
)

func main() {
	// Make an HTTP request to ServiceB
	resp, err := http.Get("http://localhost:8081/callFunction")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		helper.DataDogHandle.LogError("Error reading response body:", err)
		return
	}

	// Print the response
	fmt.Println(string(body))
}
