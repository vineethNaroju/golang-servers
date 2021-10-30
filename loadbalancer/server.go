package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer(port string) {
	x := http.NewServeMux()

	x.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from " + port)
	})

	log.Fatal(http.ListenAndServe(":" + port, x))
}