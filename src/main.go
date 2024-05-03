package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/manish-neemnarayan/go-native-server/types"
)

func main() {
	fmt.Println("server started")
	svc := NewService()

	http.HandleFunc("/health", HealthHandler())
	http.HandleFunc("/insert", InsertHandler(svc))
	http.HandleFunc("/get", GetHandler(svc))
	log.Fatal(http.ListenAndServe(":3003", nil))
}

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"data": "Server is running"})
	}
}

func InsertHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data types.PostData
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Fatal("Bad request error")
		}

		if err := svc.Insert(&data); err != nil {
			log.Fatal("Internal Server Error")
		}

		w.WriteHeader(200)
		fmt.Fprintf(w, "inserted successfully")
	}
}

func GetHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query()["id"][0]
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Bad Request Error")
		}

		data, err := svc.Get(intId)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v", *data)
		w.WriteHeader(200)
	}
}
