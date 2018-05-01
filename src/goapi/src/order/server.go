package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"util"
)

type OrderServer struct {
	om         *OrderManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer(configFile string) *OrderServer {
	n := negroni.Classic()
	addrs := util.GetAddrs(configFile)
	orderServer := &OrderServer{
		om:         NewOrderManager(addrs),
		httpServer: n}
	log.Println("Create OrderServer")
	return orderServer
}

func (os *OrderServer) Init() {
	mx := mux.NewRouter()
	os.initRouteTable(mx)
	os.httpServer.UseHandler(mx)
	log.Println("Init HTTP Request Route")
}

func (os *OrderServer) Run() {
	os.httpServer.Run()
}

func (os *OrderServer) initRouteTable(mx *mux.Router) {
	mx.HandleFunc("/users/{userid}/order", os.createOrder).Methods("POST")
	mx.HandleFunc("/users/{userid}/order/{orderid}", os.getOrder).Methods("GET")
	mx.HandleFunc("/users/{userid}/order/{orderid}", os.updateOrder).Methods("POST")
	mx.HandleFunc("/users/{userid}/order/{orderid}", os.deleteOrder).Methods("DELETE")
	mx.HandleFunc("/users/{userid}/orders", os.getOrderByUser).Methods("GET")
}

func (os *OrderServer) getOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	orderId := params["orderid"]
	val, ok := os.om.GetOrder(userId, orderId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Order not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
	log.Printf("GET Order %v\n", ok)
}

func (os *OrderServer) createOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing UserId"))
		return
	}

	items := strings.Split(req.FormValue("items"), ",")
	val, ok := os.om.CreateOrder(userId, items)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Order create failed"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
	log.Printf("CREATE Order %v\n", ok)
}

func (os *OrderServer) updateOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	orderId := params["orderid"]
	var order Order
	orderJson, ok := os.om.GetOrder(userId, orderId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal([]byte(orderJson), &order)
	addItems := strings.Split(req.FormValue("add"), ",")
	delItems := strings.Split(req.FormValue("delete"), ",")

	items := make(map[string]string)
	for _, item := range order.Items {
		items[item] = item
	}

	for _, del := range delItems {
		delete(items, del)
	}

	for _, add := range addItems {
		items[add] = add
	}

	order.Items = order.Items[:0]
	for k, _ := range items {
		order.Items = append(order.Items, k)
	}
	ok = os.om.UpdateOrder(&order)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update order failed"))
	} else {
		buf, _ := json.Marshal(order)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
	log.Printf("UPDATE Order %v\n", ok)
}

func (os *OrderServer) deleteOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	orderId := params["orderid"]
	ok := os.om.DeleteOrder(userId, orderId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete order failed"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Delete order Successfully"))
	}
	log.Printf("DELETE Order %v\n", ok)
}

func (os *OrderServer) getOrderByUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Parameter"))
		return
	}

	orders, ok := os.om.GetOrderByUser(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Get order by user failed"))
	} else {
		buf, _ := json.Marshal(orders)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
	log.Printf("GET_ORDER_BY_USER Order %v\n", ok)
}
