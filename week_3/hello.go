package main

import (
	"log"
	"net/http"
)



func helloGoHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello net/http\n"))
}

func main(){
	http.HandleFunc("/", helloGoHandler)
	log.Fatal(http.ListenandServe(":9090", nil))

}
