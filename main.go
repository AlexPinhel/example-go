package main

import (
	"fmt"
	"flag"
	"log"
	"encoding/json"
  "io"
  "net/http"
)

type Response struct {
	Hello string
}

func main() {
	fmt.Println("Hello, alex")

	var addr string

	flag.StringVar(&addr, "addr", "127.0.0.1:8000", "Bind to this address:port")
	flag.Parse()
	http.HandleFunc("/", HelloWorld)
  err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  var response Response
	response.Hello = "Good evening World"
	b, err := json.Marshal(response)
	    if err != nil {
	        http.Error(w, err.Error(), 500)
	        return
	    }
	    io.WriteString(w, string(b))
}
