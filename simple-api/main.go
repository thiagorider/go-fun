package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/beevik/ntp"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func infoPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Informations on your Request\n\n\n")
	data, err := httputil.DumpRequest(r, false)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Fprintf(w, "%s", string(data))
	fmt.Printf("%s", string(data))
	fmt.Println("Endpoint Hit: infoPage")
}

func whatTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whats time is it?")
	ntpTime, err := ntp.Query("pool.ntp.org")
	if err != nil {
		fmt.Println(err)
	}

	ntpTimeFormatted := ntpTime

	fmt.Printf("Network time: %v\n", ntpTime)
	fmt.Printf("Unix Date Network time: %v\n", ntpTimeFormatted)
	fmt.Println("+++++++++++++++++++++++++++++++")
	timeFormatted := time.Now().Local().Format(time.UnixDate)
	fmt.Printf("System time: %v\n", time.Now())
	fmt.Printf("Unix Date System time: %v\n", timeFormatted)
	//fmt.Fprintf(w, "%t", string(time))
	//fmt.Printf("%t", string(time))
	fmt.Println("Endpoint Hit: infoPage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/info", infoPage)
	http.HandleFunc("/getTime", whatTime)
	log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
	handleRequests()
}
