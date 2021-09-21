package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	jsonFile, err := ioutil.ReadFile("quotes.json")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("Successfully opened quotes.json")
	jsonString := string(jsonFile)
	fmt.Println(jsonString)
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


func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
