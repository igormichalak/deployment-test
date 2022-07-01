package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DefaultServerPort = 8080

func main() {
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if serverPort == "" || !ok {
		serverPort = strconv.Itoa(DefaultServerPort)
	}
	fmt.Println(serverPort)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The portfolio is currently under development.")
	})
	log.Fatal(http.ListenAndServe(":"+serverPort, handler))
}
