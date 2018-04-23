package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var redis_server_ip = "127.0.0.1"
var redis_server_port = 6379

type ReviewServer struct {
	om         *ReviewManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer() *ReviewServer {
	n := negroni.Classic()
	reviewServer := &ReviewServer{
		om:         NewReviewManager(redis_server_ip, redis_server_port),
		httpServer: n}
	log.Println("Create ReviewServer")
	return reviewServer
}

func (os *ReviewServer) Init() {
	mx := mux.NewRouter()
	os.initRouteTable(mx)
	os.httpServer.UseHandler(mx)
	log.Println("Init HTTP Request Route")
}

func (os *ReviewServer) Run() {
	os.httpServer.Run()
}

func (os *ReviewServer) initRouteTable(mx *mux.Router) {
	mx.HandleFunc("/review", os.createReview).Methods("POST")
	mx.HandleFunc("/review/{reviewid}", os.getReview).Methods("GET")
	mx.HandleFunc("/review/{reviewid}", os.updateReview).Methods("POST")
	mx.HandleFunc("/review/{reviewid}", os.deleteReview).Methods("DELETE")
	mx.HandleFunc("/review", os.getReviewByUser).Methods("GET")
}

func (os *ReviewServer) getReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	reviewId := params["reviewid"]
	val, ok := os.om.GetReview(reviewId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Review not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
	log.Printf("GET Review %v\n", ok)
}

func (os *ReviewServer) createReview(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("reviewid")
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter"))
		return
	}

	item := req.FormValue("item")
	val, ok := os.om.CreateReview(userId, item)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Review not exist"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
	log.Printf("CREATE Review %v\n", ok)
}

func (os *ReviewServer) updateReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	reviewid := params["reviewid"]
	var review Review
	reviewJson, ok := os.om.GetReview(reviewid)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal([]byte(reviewJson), &review)
	addItems := strings.Split(req.FormValue("add"), ",")
	delItems := strings.Split(req.FormValue("delete"), ",")

	items := make(map[string]string)
	item := review.Item
	items[item] = item

	for _, del := range delItems {
		delete(items, del)
	}

	for _, add := range addItems {
		items[add] = add
	}

	review.Item = review.Item
	for k, _ := range items {
		review.Items = append(review.Items, k)
	}
	ok = os.om.UpdateReview(&review)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Update review failed"))
	} else {
		buf, _ := json.Marshal(review)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
	log.Printf("UPDATE Review %v\n", ok)
}

func (os *ReviewServer) deleteReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	reviewId := params["reviewid"]
	ok := os.om.DeleteReview(reviewId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete review failed"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Delete review Successfully"))
	}
	log.Printf("DELETE Review %v\n", ok)
}

func (os *ReviewServer) getReviewByUser(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("reviewid")
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Parameter"))
		return
	}

	reviews, ok := os.om.GetReviewByInventory(userId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Get review by user failed"))
	} else {
		buf, _ := json.Marshal(reviews)
		w.WriteHeader(http.StatusOK)
		w.Write(buf)
	}
	log.Printf("GET_Reivew_BY_USER Review %v\n", ok)
}
