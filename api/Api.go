package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ListenAndServe() {
	http.HandleFunc("/todo", handleTodo)
	http.HandleFunc("/files", handleFiles)

	port := ":3333"
	fmt.Println("Server available on: http://localhost" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	todos := []string{
		"Parse request body (JSON).",
		"Get file path from request body.",
		"Return file content.",
	}

	fmtTodos(todos)
	res := stringsToBytes(todos, "\n\n")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func fmtTodos(todos []string) {
	for i := range todos {
		if i == 0 {
			todos[i] = "NEXT: " + todos[i]
		} else {
			todos[i] = "      " + todos[i]
		}
	}
}

func handleFiles(w http.ResponseWriter, r *http.Request) {
	panic("Not yet implemented. See /todo")
}

func stringsToBytes(v []string, delim string) []byte {
	joined := strings.Join(v, delim)
	return []byte(joined)
}
