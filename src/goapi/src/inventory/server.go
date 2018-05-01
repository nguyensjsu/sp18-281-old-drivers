package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"util"
)

type InventoryServer struct {
	im         *InventoryManager
	httpServer *negroni.Negroni
}

func NewServer(configFile string) *InventoryServer {
	n := negroni.Classic()
	addrs := util.GetAddrs(configFile)
	inventoryServer := &InventoryServer{
		im:         NewInventoryManager(addrs),
		httpServer: n}
	log.Println("Create InventoryServer")
	return inventoryServer
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
func (is *InventoryServer) initRouteTable(mx *mux.Router) {
	// mx.HandleFunc("/inventorys", is.getAllInventoryHandler).Methods("GET")
	mx.HandleFunc("/inventory/{inventoryid}", is.getInventoryHandler).Methods("GET")
	mx.HandleFunc("/inventory", is.addInventoryHandler).Methods("POST")
	mx.HandleFunc("/inventory/{inventoryid}", is.updateInventoryHandler).Methods("PUT")
	mx.HandleFunc("/inventory/{inventoryid}", is.deleteInventoryHandler).Methods("DELETE")
}

/* API Get All Inventorys
func (is *InventoryServer) getAllInventoryHandler(w http.ResponseWriter, r *http.Request) {

	var inventory Inventory
	inventoryJson, ok := is.im.GetAllInventory()

	var objmap map[string]*json.RawMessage
	json.Unmarshal(inventoryJson, &objmap)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(inventory))
	}

	log.Printf("GET All Inventory %v\n", ok)

}
*/

// API Get Inventory
// curl -i -X GET "localhost:8080/inventory/{inventoryid}"
func (is *InventoryServer) getInventoryHandler(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	inventoryId := param["inventoryid"]
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
// curl -i -X POST "localhost:8080/inventory?name=latin&price=20&amount=100"
func (is *InventoryServer) addInventoryHandler(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	price, err1 := strconv.ParseFloat(r.FormValue("price"), 64)
	amount, err2 := strconv.ParseInt(r.FormValue("amount"), 10, 64)

	if name == "" || err1 != nil || err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameters"))
	}

	val, ok := is.im.CreateInventory(name, price, amount)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Inventory create failed"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
}

// API Update Inventory
// curl -i -X PUT "localhost:8080/inventory/7fc439c2-8dd0-4829-bb3c-64f88f03fc86?name=latin&price=20&amount=99"
func (is *InventoryServer) updateInventoryHandler(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	inventoryId := param["inventoryid"]

	var inventory Inventory
	inventoryJson, ok := is.im.GetInventory(inventoryId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory Id doesn't exist"))
	}

	json.Unmarshal([]byte(inventoryJson), &inventory)
	// to be continue...
	name := r.FormValue("name")
	price := r.FormValue("price")
	amount := r.FormValue("amount")
	inventory.InventoryName = name
	inventory.InventoryPrice, _ = strconv.ParseFloat(price, 64)
	inventory.Amount, _ = strconv.ParseInt(amount, 10, 64)

	good := is.im.UpdateInventory(&inventory)
	if !good {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update inventory failed"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory updated!"))
	}
}

// API Delete Inventory
// curl -i -X DELETE "localhost:8080/inventory/dad0c14f-e0ae-4fa3-9a8c-29b9dad9347e"
func (is *InventoryServer) deleteInventoryHandler(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	inventoryId := param["inventoryid"]
	ok := is.im.DeleteInventory(inventoryId)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Inventory not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory delete successful"))
	}

	log.Printf("DELETE Inventory %v\n", ok)

}
