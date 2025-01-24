package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

func main() {

	mux := http.DefaultServeMux //ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns,
	// and calls the handler for the pattern that most closely matches the URL.

	mux.HandleFunc("GET /posts", listPostsHandler) // HandleFunc registers the handler function for the given pattern in DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
	http.HandleFunc("POST /posts/{id}", getPostHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux)) // ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error reading request body", http.StatusBadRequest)
	// 	return
	// }
	// var post Post
	// err = json.Unmarshal(body, &post)
	var post Post
	// err = json.Unmarshal(body, &post)
	// if err != nil {
	// 	http.Error(w, "Error unmarshalling request body again", http.StatusBadRequest)
	// 	return
	// }
	post.ID = 1
	post.Name = "harsh"
	// handleResponse(w, post, http.StatusOK)
	handleResponse(w, post)
}

func listPostsHandler(w http.ResponseWriter, r *http.Request) {

	posts := []Post{
		{ID: 1, Name: "First post"},
		{ID: 2, Name: "Second post"},
	}

	// handleResponse(w, posts, http.StatusOK)
	handleResponse(w, posts)

}

// func handleResponse(w http.ResponseWriter, message any, statusCode int) {
func handleResponse(w http.ResponseWriter, message any) {
	response, _ := json.Marshal(message)
	// if err != nil {
	// 	http.Error(w, "Error converting response to json", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(statusCode)
	w.Write(response)
}
