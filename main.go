package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init(){
	file, err := ioutil.ReadFile("banner.txt")
	if err != nil{
		fmt.Println("FAILED TO READ BANNER, PROGRAMM STARTED!")
	} else{
		fileText := string(file)
		fmt.Println(fileText)
	}
}

func readJSON(){
	jsonFile, err := os.Open("quotes.json")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("Successfully opened quotes.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quotes []Quote
	
}

func main(){
	fmt.Println("Hello World!")
	readJSON()
	handleRequests()
}

func homePage(w http.ResponseWriter, r*http.Request){
	_, _ = fmt.Fprintf(w, "Welocome to the homepage!")
	fmt.Println("Endpoint Hit: Homepage")
}

func randomQuote(w http.ResponseWriter, r*http.Request){
	fmt.Println("Endpoint Hit: randomQuote")
}


func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
