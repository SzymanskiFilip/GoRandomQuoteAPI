package main

import (
	"fmt"
	"io/ioutil"
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


func main(){
	fmt.Println("Hello World!")

}
