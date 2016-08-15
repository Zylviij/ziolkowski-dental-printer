package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "http://ziolkowskidental.com/info.php", nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("User-Agent", "this is a very secret password")
	request.SetBasicAuth("Usrnm", "Psswrd")
	fmt.Println(request)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", body)
}
