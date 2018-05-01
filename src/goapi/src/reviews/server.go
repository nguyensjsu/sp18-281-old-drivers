package main

import (
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"util"
)

type ReviewServer struct {
	rm         *ReviewManager
	httpServer *negroni.Negroni
}

// NewServer configures and returns a Server.
func NewServer(configFile string) *ReviewServer {
	n := negroni.Classic()
	addrs := util.GetAddrs(configFile)
	reviewServer := &ReviewServer{
		rm:         NewReviewManager(addrs),
		httpServer: n}
	log.Println("Create ReviewServer")
	return reviewServer
}

func (rs *ReviewServer) Init() {
	mx := mux.NewRouter()
	rs.initRouteTable(mx)
	rs.httpServer.UseHandler(mx)
	log.Println("Init HTTP Request Route")
}

func (rs *ReviewServer) Run() {
	rs.httpServer.Run()
}

func (rs *ReviewServer) initRouteTable(mx *mux.Router) {
	mx.HandleFunc("/users/{userid}/review", rs.createReview).Methods("POST")
	mx.HandleFunc("/users/{userid}/review/{reviewid}", rs.getReview).Methods("GET")
	mx.HandleFunc("/users/{userid}/review/{reviewid}", rs.updateReview).Methods("POST")
	mx.HandleFunc("/users/{userid}/review/{reviewid}", rs.deleteReview).Methods("DELETE")
	mx.HandleFunc("/users/{userid}/reviews", rs.getReviewByUser).Methods("GET")
}

func (rs *ReviewServer) getReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	reviewId := params["reviewid"]
	val, ok := rs.rm.GetReview(userId, reviewId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Review not exist"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	}
	log.Printf("GET Review %v\n", ok)
}

func (rs *ReviewServer) createReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter"))
		return
	}

	item := req.FormValue("item")
	content := req.FormValue("content")

	val, ok := rs.rm.CreateReview(userId, item, content)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Review not exist"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(val))
	}
	log.Printf("CREATE Review %v\n", ok)
}

func (rs *ReviewServer) updateReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	reviewid := params["reviewid"]
	newContent := req.FormValue("content")

	var review Review
	reviewJson, ok := rs.rm.GetReview(userId, reviewid)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal([]byte(reviewJson), &review)

	ok = rs.rm.UpdateReview(&review, newContent)
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

func (rs *ReviewServer) deleteReview(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	reviewId := params["reviewid"]
	ok := rs.rm.DeleteReview(userId, reviewId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Delete review failed"))
	} else {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte("Delete review Successfully"))
	}
	log.Printf("DELETE Review %v\n", ok)
}

func (rs *ReviewServer) getReviewByUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId := params["userid"]
	if len(userId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Parameter"))
		return
	}

	reviews, ok := rs.rm.GetReviewByReviewId(userId)
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
