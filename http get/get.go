package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		res, err := http.Get("http://127.0.0.1/status.txt")
		if err == nil {
			// log.Fatal(err)
			robots, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)
		}

		if err != nil {
			log.Println(err)

		}
		log.Println("Sleeping...")
		time.Sleep(5 * time.Second)
	}

}
