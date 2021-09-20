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
	defer func(jsonFile *os.File) {
		_ = jsonFile.Close()
	}(jsonFile)

	var quote Quote
	json.Unmarshal([]byte(jsonFile), &quote)
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
