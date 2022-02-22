package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Vegetable struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Shop struct {
	mp map[string]Vegetable
}

func (shop *Shop) get(name string) (*Vegetable, error) {
	if veg, has := shop.mp[name]; has {
		return &veg, nil
	} else {
		return nil, fmt.Errorf("Vegetable with name %s is not found", name)
	}
}

func (shop *Shop) put(veg Vegetable) {
	shop.mp[veg.Name] = veg
}

func (shop *Shop) delete(name string) {
	delete(shop.mp, name)
}

func main() {

	shop := &Shop{
		mp: make(map[string]Vegetable),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/get", func(rw http.ResponseWriter, req *http.Request) {
		queryparam := req.URL.Query()

		vegName, have := queryparam["name"]

		if !have {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Missing name query param"))
			return
		}

		veg, err := shop.get(vegName[0])

		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(err.Error()))
			return
		}

		rw.WriteHeader(http.StatusOK)

		json.NewEncoder(rw).Encode(veg)
	})

	mux.HandleFunc("/post", func(rw http.ResponseWriter, req *http.Request) {
		var veg Vegetable

		if err := json.NewDecoder(req.Body).Decode(&veg); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(err.Error()))
			return
		}

		shop.put(veg)

		rw.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/delete", func(rw http.ResponseWriter, req *http.Request) {
		queryparam := req.URL.Query()

		name, have := queryparam["name"]

		if !have {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Missing name query param"))
			return
		}

		shop.delete(name[0])

		rw.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:5555", mux))
}
