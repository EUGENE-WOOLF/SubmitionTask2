package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// "encoding/json"
var StudentList []StudentInfo

func SubmitHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	//fixing the COROS issue and now dealing with the acual request

	if r.Method != http.MethodPost {
		fmt.Println("There was error html forms")
		return
	}

	var student StudentInfo

	err := json.NewDecoder(r.Body).Decode(&student)
	defer r.Body.Close()

	if err != nil {
		fmt.Println("Error in json file")
		return
	}
	fmt.Printf("Received: %+v\n %T", student, student)

	StudentList = append(StudentList, student)
	MakingListOfStudents(student)

	w.WriteHeader(http.StatusOK)

}
