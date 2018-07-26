package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://api.theysaidso.com/qod.json")
	checkError(err)

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	checkError(err)

	fmt.Println(string(contents))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
