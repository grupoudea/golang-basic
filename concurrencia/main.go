package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://api.fakejg.com",
		"https://graph.microsoft.com",
	}

	channel := make(chan string)

	for _, api := range apis {
		go checkApi(api, channel)
		//checkApiNoConcurrent(api)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-channel)
	}

	end := time.Since(start)

	fmt.Printf("Tardó %v", end)

}

func checkApi(api string, channel chan string) {
	_, err := http.Get(api)

	if err != nil {
		channel <- fmt.Sprintln("Error: ", err)
		return
	}

	channel <- fmt.Sprintf("Funciona la api %s\n", api)

}

func checkApiNoConcurrent(api string) {
	_, err := http.Get(api)

	if err != nil {
		fmt.Printf("Error: %s - %s\n", api, err)
		return
	}

	fmt.Printf("Funciona la api %s\n", api)
}
