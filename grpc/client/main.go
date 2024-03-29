package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address = "grpc-server:4000"
	//address = "localhost:4000"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func newElement(w http.ResponseWriter, r *http.Request) {
	// HEADERS
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"message\": \"ok\"}"))
		return
	}

	//Parsing body
	var body map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	failOnError(err, "Parsing JSON")
	body["way"] = "GRPC"
	data, err := json.Marshal(body)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	failOnError(err, "GRPC Connection")
	defer conn.Close()

	//Adding new client
	cli := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	re, err := cli.SayHello(ctx, &pb.HelloRequest{Name: string(data)})

	if err != nil {
		failOnError(err, "Error al enviar el mensaje")
	}
	log.Print("Sent:")
	log.Printf("Response : %s", re.GetMessage())

	// Setting status and response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(re.GetMessage()))
}

func handleRequest() {
	http.HandleFunc("/", newElement)
	log.Print("Client listenin on port: 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	handleRequest()
}
