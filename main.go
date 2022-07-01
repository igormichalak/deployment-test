package main

import (
	"fmt"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

//const DefaultServerPort = 8080

func main() {
	//serverPort, ok := os.LookupEnv("SERVER_PORT")
	//if serverPort == "" || !ok {
	//	serverPort = strconv.Itoa(DefaultServerPort)
	//}
	//fmt.Println("Server running on port " + serverPort)

	m := &autocert.Manager{
		Cache:      autocert.DirCache("secret-dir"),
		Prompt:     autocert.AcceptTOS,
		Email:      "hello@igormichalak.com",
		HostPolicy: autocert.HostWhitelist("igormichalak.com", "www.igormichalak.com", "api.igormichalak.com"),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "The portfolio is currently under development.")
	})

	s := &http.Server{
		Addr:      ":443",
		Handler:   handler,
		TLSConfig: m.TLSConfig(),
	}

	go http.ListenAndServe(":80", m.HTTPHandler(nil))
	log.Fatal(s.ListenAndServeTLS("", ""))
}
