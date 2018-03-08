package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		log.Fatal(err)
	}
	freebook, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", freebook)
}
