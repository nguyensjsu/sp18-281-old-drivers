package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/url"
	// "github.com/satori/go.uuid"
	// "github.com/unrolled/render"
)

func NewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"",
		DB: 		1,
		})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/inventorys", getInventorysHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory/{id}", getInventoryHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory", addInventoryHandler(formatter)).Methods("POST")
	mx.HandleFunc("/inventory/{id}", updateInventoryHandler(formatter)).Methods("PUT")
	mx.HandleFunc("/inventory/{id}", gumballNewOrderHandler(formatter)).Methods("DELETE")
}


// API Get Inventorys
func getInventorysHandler(formatter *render.Render) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		var cursor uint64
		x := make(map[string][])
		for i := 0; i < client.; i++ {
			
		}
	}
}

// API Get Inventory
func getInventoryHandler(formatter *render.Render) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r, err := url.ParseForm(r)
		if err != nil {
			panic(err)
		}
		if len(r.Form["id"] > 0) {
			formatter.JSON(w, http.StatusOK, client.HScan(cursor, r.FormValue["id"], ))
		}
	}
}

// API Add Inventory
func addInventoryHandler(formatter *render.Render) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r, err := url.ParseForm(r)
		if err != nil {
			panic(err)
		}

		uuid := uuid.NewV4()
		client.HMset(uuid, inventoryId r.FormValue)



		/*
		client {
			inventoryId:	uuid.String(),
			inventoryName:	r.FormValue("name"),
			inventoryPrice: r.FormValue("price"),
			inventoryLeft:  r.FormValue("amount"),
		}
		*/

		


		
	}
}

// API Update Inventory
func updateInventoryHandler(formatter *render.Render) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}

// API Delete Inventory
func deleteInventoryHandler(formatter *render.Render) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}
