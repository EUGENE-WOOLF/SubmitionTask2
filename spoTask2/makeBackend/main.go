package main

import (
	"fmt"
	"log"
	"net/http"
)

//how it works
//form submites data ---> json.NewDecode --> gets json --> use for loop(also has a counter) --> adds to a slice --> write from a sclice

func main() {

	http.HandleFunc("/submit", SubmitHandler)
	http.HandleFunc("/summary", SummaryHandler)
	SettingUpHeaderCsvFile()
	fmt.Println("Server running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
