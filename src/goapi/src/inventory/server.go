package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net/url"
	"github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

var redis_server_ip = "127.0.0.1"
var redis_server_port = 6379

type InventoryServer struct {
	im         *InventoryManager
	httpServer *negroni.Negroni
}

func NewServer() *InventoryServer {
	n := negroni.Classic()
	inventoryServer := &InventoryServer{
		im:         InventoryManager(redis_server_ip, redis_server_port),
		httpServer: n}
	log.Println("Create InventoryServer")
	return orderServer
}

func (is *InventoryServer) Init() {
	mx := mux.NewRouter()
	is.initRouteTable(mx)
	is.httpServer.UseHandler(mx)
	log.Println("Init HTTP Request Route")
}

func (is *InventoryServer) Run() {
	is.httpServer.Run()
}

// API Routes
func (is *InventoryServer) initRoutesTable(mx *mux.Router) {
	mx.HandleFunc("/inventorys", is.getInventorysHandler()).Methods("GET")
	mx.HandleFunc("/inventory/{id}", is.getInventoryHandler()).Methods("GET")
	mx.HandleFunc("/inventory", is.addInventoryHandler()).Methods("POST")
	mx.HandleFunc("/inventory/{id}", is.updateInventoryHandler()).Methods("PUT")
	mx.HandleFunc("/inventory/{id}", is.gumballNewOrderHandler()).Methods("DELETE")
}


// API Get All Inventorys
func (is *InventoryServer) getInventorysHandler(w http.ResponseWriter, r *http.Request) {
	
	val, ok := is.im.GetAllInventory()

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}

	log.Printf("GET All Inventory %v\n", ok)

}

// API Get Inventory
func (is *InventoryServer) getInventoryHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Var(r)
	inventoryId := param["inventoryId"]
	val, ok := is.im.GetInventory(inventoryId)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}

	log.Printf("GET Inventory %v\n", ok)

}

// API Add Inventory
func (is *InventoryServer) addInventoryHandler(w http.ResponseWriter, r *http.Request) {
	
	name := r.FormValue("name")
	price := float32(r.FormValue("price"))
	amount := uint32(r.FormValue("amount"))

	if inventoryName == nil || inventoryPrice == nil || amount == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader("Invalid parameters")
	}
	
	val, ok := is.im.CreateInventory(name, price, amount)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader("Inventory create failed")
	}
	else {
		w.WriteHeader(http.StatusCreated)
		w.WriteHeader("Inventory created!")
	}
}


// API Update Inventory
func (is *InventoryServer) updateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	
	param := mux.Var(r)
	inventoryId := param["inventoryId"]

	var inventory Inventory
	inventoryJson, ok := is.io.GetInventory(inventoryId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.WriteHeader("Inventory Id doesn't exist")
	}

	json.Unmarshal([]byte(inventoryJson), &inventory)
	// to be continue...
	name := r.FormValue["name"]
	price := r.FormValue["price"]
	amount := r.FormValue["amount"]
	inventory.InventoryName = name
	inventory.InventoryPrice = price
	inventory.Amount = amount

	val, ok := is.im.UpdateInventory(inventory)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.WriteHeader("Update inventory failed")
	}
	else {
		w.WriteHeader(http.StatusOK)
		w.WriteHeader("Inventory updated!")
	}
}

// API Delete Inventory
func (is *InventoryServer) deleteInventoryHandler(w http.ResponseWriter, r *http.Request) {
	
	param := mux.Var(r)
	inventoryId := param["inventoryId"]
	val, ok := is.im.DeleteInventory(inventoryId)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}

	log.Printf("DELETE Inventory %v\n", ok)	
	
}
