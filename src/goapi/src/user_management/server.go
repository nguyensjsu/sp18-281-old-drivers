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

type UserServer struct {
	um         *UserManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer(configFile string) *UserServer {
	n := negroni.Classic()
	addrs := util.GetAddrs(configFile)
	UserServer := &UserServer{
		um:         NewUserManager(addrs),
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

// get user
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

// create user
func (us *UserServer) createUser(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	phone := req.FormValue("phone")
	balance := req.FormValue("balance")
	val, ok := us.um.CreateUser(name, phone, balance)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User create failed"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
	log.Printf("CREATE User %v\n", ok)
}

// update user
func (us *UserServer) updateUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	userJson, ok := us.um.GetUser(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var user User
	json.Unmarshal([]byte(userJson), &user)
	phone := req.FormValue("phone")
	balance := req.FormValue("balance")
	if len(phone) != 0 {
		user.Phone = phone
	}
	if len(balance) != 0 {
		bal, err := strconv.Atoi(balance)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user.Balance = bal
	}

	ok = us.um.UpdateUser(&user)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update user failed"))
	} else {
		buf, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
	log.Printf("UPDATE User %v\n", ok)
}

// delete user
func (us *UserServer) deleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	ok := us.um.DeleteUser(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete user failed"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Delete user Successfully"))
	}
	log.Printf("DELETE user %v\n", ok)
}
