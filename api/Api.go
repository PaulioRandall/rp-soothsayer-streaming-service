package api

import (
	"fmt"
	"log"
	"net/http"
)

func ListenAndServe() {
	attachHttpFileListHandler()
	attachHttpFileHandler()

	port := ":3333"
	fmt.Println("Server available on: http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func appendCorsHeaders(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
