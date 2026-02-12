package main

import (
	"net/http"
	"sync"
)

type Sites struct {
	url    string
	status int
}

func main() {
	sitesUrl := []Sites{
		{url: "www.facebook.com", status: 0},
		{url: "www.google.com", status: 0},
		{url: "www.youtube.com", status: 0},
	}
	stringsChannel := make(chan string, 100)
	sitesChannel := make(chan Sites, 100)

	var waitinGroup sync.WaitGroup

	for i := 0; i < 3; i++ {
		stringsChannel <- sitesUrl[i].url
		waitinGroup.Add(1)
		go worker(stringsChannel, sitesChannel, &waitinGroup)
	}

	waitinGroup.Wait()
	close(stringsChannel)

}

func worker(jobs chan string, results chan Sites, wg *sync.WaitGroup) {
	defer wg.Done()
	for trabajo := range jobs {
		//Aqui procesaremos las URL
		resp, err := http.Get(trabajo)
		if err != nil {
			results <- Sites{url: trabajo, status: 0}
			continue
		}

		results <- Sites{url: trabajo, status: resp.StatusCode}

	}
}
