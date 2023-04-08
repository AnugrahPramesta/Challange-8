package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for _ = range time.Tick(time.Second * 15) {
		wind := rand.Intn(20)
		water := rand.Intn(20)
		data := map[string]interface{}{
			"wind":  wind,
			"water": water,
		}
		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		// defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(body))

		if wind < 6 {
			fmt.Println("status wind : Aman")
		} else if wind >= 7 && wind <= 15 {
			fmt.Println("status wind : Siaga")
		} else if wind > 15 {
			fmt.Println("status wind : Bahaya")
		}
		if water < 5 {
			fmt.Println("status water : Aman")
		} else if water >= 6 && water <= 8 {
			fmt.Println("status water : Siaga")
		} else if water > 8 {
			fmt.Println("status water : Bahaya")
		}
		// fmt.Println(data)
	}
}
