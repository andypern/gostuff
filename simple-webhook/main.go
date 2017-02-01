package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"



)

//thanks daniel for this...

var (
	listenPort       = flag.String("listenPort", "8087", "port to listen on")

)




func main() {
	flag.Parse()




	log.Fatal(http.ListenAndServe(":"+*listenPort,
	 http.HandlerFunc(
	  func(w http.ResponseWriter, r *http.Request) {
	    b, _ := ioutil.ReadAll(r.Body)

	    //fmt.Printf("Req: %#v\nBody: %s\n", r, b)
	    fmt.Printf("%s\n", b)
	    w.WriteHeader(http.StatusOK)
	})))
}
