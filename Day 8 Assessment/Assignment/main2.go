// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type StatusChecker interface {
// 	Check(ctx context.Context, name string) (status bool, err error)
// }
// type httpChecker struct {
// }

// type WebsitesHandler struct {
// 	ctx           context.Context
// 	sitesStatus   map[string]string
// 	statusChecker StatusChecker
// }

// // ServeHTTP implements http.Handler.
// func (webHandle WebsitesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// }

// func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
// 	// your implementation to check status using HTTP call

// 	return
// }

// func main() {
// 	ctx := context.Background()

// 	httpObj := httpChecker{}
// 	sitesStatus := map[string]string{}

// 	siteHandler := WebsitesHandler{
// 		ctx:           ctx,
// 		sitesStatus:   sitesStatus,
// 		statusChecker: httpObj,
// 	}

// 	mux := http.DefaultServeMux

// 	mux.Handle("GET /websites", siteHandler)
// 	mux.Handle("POST /websites", siteHandler)

// 	fmt.Println("Listening to Port :8080")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }

// func handleRequest(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		siteName := r.URL.Query().Get("name")
// 		if len(siteName) != 0 {
// 			handleSingleWebsite(ctx)
// 		} else {
// 			fmt.Println("Fuck you")
// 		}

// 	} else {
// 		fmt.Println("Fuck you")
// 	}
// }

// func handleSingleWebsite(ctx context.Context, name string) {

// }
