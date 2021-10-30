package loadbalancer

import (
	"log"
	"net/http"
)

func Test() {

	ports := []string {"8081", "8082", "8083"}

	for _, port := range ports {
		go StartServer(port)
	}

	sp := NewServerPool()

	x := http.NewServeMux()

	sp.AddServer("http://localhost:8081")
	sp.AddServer("http://localhost:8082")
	sp.AddServer("http://localhost:8083")

	x.HandleFunc("/", sp.LoadBalance)

	log.Fatal(http.ListenAndServe(":8080", x))
}