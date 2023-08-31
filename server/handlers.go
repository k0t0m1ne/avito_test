package main

import (
	"avitoTest/repository/read"
	"avitoTest/repository/write"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func userSegmentHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	u := read.GetUserById(db, id)
	json, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func addSegmentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type body struct {
		Seg string `json:"segment"`
	}
	var b body
	decoder.Decode(&b)
	write.AddSegment(db, b.Seg)
}

func deleteSegmentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	type body struct {
		Seg string `json:"segment"`
	}
	var b body
	decoder.Decode(&b)

	write.DeleteSegment(db, b.Seg)
}

func addUserToSegment(w http.ResponseWriter, r *http.Request) {
	//a, _ := io.ReadAll(r.Body)
	//fmt.Println(string(a))
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	type body struct {
		SegmentsToAdd    []string `json:"segments_to_add"`
		SegmentsToDelete []string `json:"segments_to_delete"`
	}
	var b body

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&b)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
	defer r.Body.Close()
	write.AddUserToSegment(db, b.SegmentsToAdd, b.SegmentsToDelete, id)
}
