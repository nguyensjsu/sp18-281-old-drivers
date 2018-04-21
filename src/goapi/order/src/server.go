package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strings"
)

var redis_master_name = "redis-master"
var redis_sentinel_addr = "127.0.0.1:26379"

type OrderServer struct {
	om         *OrderManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer() *OrderServer {
	n := negroni.Classic()
	orderServer := &OrderServer{
		om:         NewOrderManager(redis_master_name, redis_sentinel_addr),
		httpServer: n}
	return orderServer
}

func (os *OrderServer) Init() {
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
}

func (os *OrderServer) Run() {
	os.httpServer.Run()
}

func (os *OrderServer) initRouteTable(mx *mux.Router) {
	mx.HandleFunc("/order", createOrder).Methods("POST")
	mx.HandleFunc("/order/{orderid}", getOrder).Methods("GET")
	mx.HandleFunc("/order/{orderid}", updateOrder).Methods("POST")
	mx.HandleFunc("/order/{orderid}", deleteOrder).Methods("DELETE")
	mx.HandleFunc("/order", getOrderByUser).Methods("GET")
}

func (os *OrderServer) getOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	orderId := params["orderid"]
	val, ok := om.GetOrder(orderId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Order not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
}

func (os *OrderServer) createOrder(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter"))
		return
	}

	items := strings.Split(req.FromValue("items"), ",")
	val, ok := om.CreateOrder(userId, items)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Order not exist"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
}

func (os *OrderServer) updateOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	orderId := params["orderid"]
	var order Order
	orderJson, ok := om.GetOrder(orderId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.Unmarshal([]byte(orderJson), &order)
	addItems := strings.Split(req.FromValue("add"), ",")
	delItems := strings.Split(req.FromValue("delete"), ",")

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
	ok := om.UpdateOrder(&order)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update order failed"))
	} else {
		buf, _ := json.Marshal(order)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}

func (os *OrderServer) deleteOrder(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	orderId := params["orderid"]
	ok := om.DeleteOrder(orderId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete order failed"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Delete order Successfully"))
	}
}

func (os *OrderServer) getOrderByUser(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Parameter"))
		return
	}

	orders, ok := om.GetOrderByUser(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Get order by user failed"))
	} else {
		buf, _ := json.Marshal(orders)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
}
