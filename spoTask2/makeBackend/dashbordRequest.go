package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SummaryHandler(w http.ResponseWriter, r *http.Request) {

	//this is fixing the coros issue
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	//making sure this is a get request only
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	e := json.NewEncoder(w).Encode(StudentList)
	// var sliceOfStudetns []StudentInfo

	// for _, s := range StudentList {
	// 	e := json.NewEncoder(w).Encode(s)

	if e != nil {
		fmt.Println("There was  a bug")
	}
	// }

}
