package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Site struct {
	Url    string
	Status int
}

func (s *Site) check() {

	resp, err := http.Get(s.Url)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	s.Status = resp.StatusCode
	fmt.Println(s.Url, resp.StatusCode)

}

func mainRoutine() {
	var grupoEspera sync.WaitGroup

	sitios := []Site{
		{Url: "https://www.google.com"},
		{Url: "https://facebook.com"},
		{Url: "https://stackoverflow.com"},
	}

	fmt.Println(sitios)

	grupoEspera.Add(len(sitios))

	for i := 0; i < len(sitios); i++ {
		go func(s *Site) {
			defer grupoEspera.Done()
			s.check()
			fmt.Printf("Terminó: %s con status %d\n", s.Url, s.Status)
		}(&sitios[i])
	}

	fmt.Println("Esperando a que terminen las peticiones...")

	grupoEspera.Wait()

	fmt.Println("¡Todo el proceso finalizado!")
}
