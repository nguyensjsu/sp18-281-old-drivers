package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"os"
)

var redis_server_ip = "127.0.0.1"
var redis_server_port int = 6379

type UserServer struct {
	um         *UserManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer() *UserServer {
	n := negroni.Classic()
	UserServer := &UserServer{
		um:         NewUserManager(redis_server_ip, redis_server_port),
		httpServer: n}
	log.Println("Create UserServer")
	return UserServer
}

func (us *UserServer) Init() {
	mx := mux.NewRouter()
	us.initRouteTable(mx)
	us.httpServer.UseHandler(mx)
	log.Println("Init HTTP Request Route")
}

func (us *UserServer) Run() {
	us.httpServer.Run()
}

func (us *UserServer) initRouteTable(mx *mux.Router) {
	mx.HandleFunc("/user", us.createUser).Methods("POST")
	mx.HandleFunc("/user/{userid}", us.getUser).Methods("GET")
	mx.HandleFunc("/user/{userid}", us.updateUser).Methods("POST")
	mx.HandleFunc("/user/{userid}", us.deleteUser).Methods("DELETE")
}

func (us *UserServer) getUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	val, ok := us.um.GetUser(userId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
	log.Printf("GET User %v\n", ok)
}

func (us *UserServer) createUser(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter"))
		return
	}

	/*items := strings.Split(req.FormValue("items"), ",")
	val, ok := os.om.CreateOrder(userId, items)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Order create failed"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}*/
	log.Printf("CREATE User %v\n", ok)
}

func (us *UserServer) updateUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	//var user User
	userJson, ok := us.um.GetUser(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal([]byte(userJson), &user)
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
	orderId := params["orderid"]
	ok := os.om.DeleteOrder(orderId)
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
