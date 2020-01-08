package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckCreateRequest() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			//This function will check if a request is correct or no
			decoder := json.NewDecoder(r.Body)
			var data User
			if err := decoder.Decode(&data); err != nil {
				fmt.Fprintf(w, "error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if data.IsValid() {
				fmt.Printf("%+v\n", data)

				response, err := data.ToJson()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
				f(w, r)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Missing or Invalid Fields")
			}

		}
	}
}

func Logging() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()
			f(w, r)
		}
	}
}

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Checking Authentication")
			f(w, r)
		}
	}
}
