package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
	"os"
	"encoding/json"

	"github.com/optiopay/kafka"
	//"github.com/optiopay/kafka/proto"

)

var (
	listenPort       = flag.String("listenPort", "8087", "port to listen on")
	kafkaAddrs     = flag.String("kafkaHost", "localhost:9092 localhost:9093", "space separated host:port pairs for kafka brokers")
	topic      = flag.String("topic", "mytopic", "topic you want to produce to")
)



type kafkaHandler struct {

    kafkaClient kafka.Client
}


func main() {
	flag.Parse()


	//turn our string into a string array
	//kAddrs := strings.Fields(*kafkaAddrs)

	// Here's how we define our kafkaConnection
	/*kConnect,err := kafka.Dial(kAddrs, kafka.NewBrokerConf("test-client"))
	if err != nil {
		fmt.Printf("screwed up on broker: %v\n", err)
		os.Exit(1)
	}
*/
// create a handler that we can use
	//kHandler := &kafkaHandler{kafkaClient: kConnect}


	log.Fatal(http.ListenAndServe(":"+*listenPort,
	 http.HandlerFunc(
	  func(w http.ResponseWriter, r *http.Request) {
	    b, _ := ioutil.ReadAll(r.Body)
		jb, err := json.Marshal(b)
		if err != nil {
			fmt.Printf("error serializing event: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", jb)
		//this is where to insert call to kafka func

	    //fmt.Printf("Req: %#v\nBody: %s\n", r, string(b))
	    w.WriteHeader(http.StatusOK)
	})))
}


/*
func produceEvent(broker kafka.Client, inputEvent []byte) {
    producer := broker.Producer(kafka.NewProducerConf())
    msg := &proto.Message{Value: []byte(inputEvent)}
	fmt.Printf("I made it")
    if _, err := producer.Produce(*topic, partition, msg); err != nil {
        log.Fatalf("cannot produce message to %s:%d: %s", topic, partition, err)
		fmt.Printf("I had an error: %s", err)
    }
}
*/
