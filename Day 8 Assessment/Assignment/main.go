// Question 1
// (50 points)

// Please complete the following assignment on GIT & add you mentor as reviewer on GIT.
// Once done, please write "Done"  in the asnwer section on the LMS.

// Problem Statement
// We are going to build an application that is meant to be a sort of status checker for some common websites that exist online. The application continuously polls the status of these websites and exposes APIs to retrieve the information.

// Requirements
// The application should expose an API using which we can submit lists of websites to monitor
// ● Implement an HTTP server and expose it on a port eg: 3000
// ● Expose an API endpoint (hint: POST /websites) and accept the list of websites in the request body as an array
// ● Save the list of websites in an in-memory map.
// 1. The application should monitor the status of all submitted websites every 1 minute
// ● Create a go routine which loops over all the websites and checks if they are responding to HTTP traffic (hint: status code 200). If yes, the website can be considered as UP, otherwise DOWN. Once the status check is done for all websites, sleep for 1 minute and continue this process forever.
// ● Status checks of N websites can be done concurrently using goroutines to improve performance.
// ● The status of each website can be saved in the same in-memory map where the list was stored.
// 2. The application should expose an API using which we can see the status of all websites. It should also support passing the name of a specific website and then it should only return the status of that particular website
// ● Expose an API endpoint (hint: GET /websites) which returns an array of websites with their current status
// ● Support a query parameter in API (hint: GET
// /websites?name=www.facebook.com) which then returns the status of the given website
// Requirements (Good to have)
// 1. Use of Golang interface to check the status of the website. Today we are relying on the HTTP status code of a website to determine whether it's up or not. Tomorrow, we could use some external third-party service for the same. Hence we could write up an interface to do the status check. Currently, it will use the HTTP implementation which can be switched later to advanced implementation. Below given is a sample for your reference. However, feel free to use your own names and signature.
// 2. Use of Go packages is encouraged.
// type StatusChecker interface {
// Check(ctx context.Context, name string) (status bool, err error) }
// type httpChecker struct {
// }
// func (h httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
// // your implementation to check status using HTTP call
// return
// }
// Sample Input (for POST /websites)
// {
// "websites" : ["www.google.com","www.facebook.com",”www.fakewebsite1.com”] }
// Sample Response (for POST /websites)
// 200 OK
// Sample Input (for GET /websites)
// Query param (optional) ?name=www.facebook.com
// Sample Response (for GET /websites)
// {
// "www.facebook.com" : "UP",
// "www.google.com" : "UP",
// "www.fakewebsite1.com”: “DOWN”, }

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type requestData struct {
	Websites []string `json:"websites"`
}

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
}

type httpChecker struct {
	ctx         context.Context
	sitesStatus map[string]bool
}

func (h httpChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// we find what kind of request it is here... so basically if it is a get request then we StatusCheck for the website passed in param
	// else if its a post request that means that a list of sites are passed, hence we check for all of them using a different function maybe.. for Granularity or Modularity

	method := r.Method

	if method == "GET" {
		websiteName := r.URL.Query().Get("name")

		if len(websiteName) == 0 {

			sitesString := []string{
				"www.facebook.com",
				"www.google.com",
				"www.fakewebsite1.com"}

			response, marshalErr := json.Marshal(sitesString)
			if marshalErr != nil {
				http.Error(w, "Error: "+marshalErr.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(response)
			return
		}

		isUp, err := h.CheckStatusOne(h.ctx, websiteName)

		if err != nil || !isUp {
			w.Write([]byte("DOWN"))
			return
		}
		w.Write([]byte("UP"))

	} else {

		for {
			body, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "Problem reading request body: "+err.Error(), http.StatusInternalServerError)
			}

			var resp requestData

			newErr := json.Unmarshal(body, &resp)

			if newErr != nil {
				http.Error(w, "Found Error: "+newErr.Error(), http.StatusInternalServerError)

			}

			h.checkStatusAll(h.ctx, resp.Websites)

			response, marshalErr := json.Marshal(h.sitesStatus)
			if marshalErr != nil {
				http.Error(w, "Error: "+marshalErr.Error(), http.StatusInternalServerError)
			}

			w.Write(response)
			time.Sleep(time.Second * 5)
		}
	}
}

func (h *httpChecker) checkStatus(ctx context.Context, site string) {
	status, err := h.CheckStatusOne(ctx, site)
	if err != nil {
		h.sitesStatus[site] = false
	}
	if status {
		h.sitesStatus[site] = true
	}
}

func (h httpChecker) checkStatusAll(ctx context.Context, siteNames []string) {
	for _, site := range siteNames {
		go h.checkStatus(ctx, site)
	}
}

func (h httpChecker) CheckStatusOne(ctx context.Context, name string) (status bool, err error) {
	resp, err := http.Get(name)
	if err != nil {
		return false, errors.New(err.Error())
	}

	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, errors.New("status code doesn't match")
}

func main() {

	mux := http.DefaultServeMux
	ctx := context.Background()

	sitesStatus := make(map[string]bool)

	httpObj := httpChecker{
		ctx,
		sitesStatus,
	}

	// mux.HandleFunc("GET /helloServer", handleGetRequest) // just for testing
	// mux.HandleFunc("POST /websites", websiteStatusHandler)

	mux.Handle("GET /websites", httpObj)
	mux.Handle("POST /websites", httpObj)

	fmt.Println("Listen to Port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// func handleGetRequest(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("Hello"))

// 	if err != nil {
// 		http.Error(w, "Error in write: "+err.Error(), http.StatusInternalServerError)
// 	}
// }

// func websiteStatusHandler(w http.ResponseWriter, r *http.Request) {

// 	body, err := io.ReadAll(r.Body)

// 	if err != nil {
// 		http.Error(w, "Problem reading request body: "+err.Error(), http.StatusInternalServerError)
// 	}

// 	var data requestData

// 	errr := json.Unmarshal(body, &data)

// 	if errr != nil {
// 		http.Error(w, "Error found: "+errr.Error(), http.StatusInternalServerError)
// 	}

// 	websiteName := r.URL.Query().Get("name")

// 	// if no params passed check for all sites

// 	if len(websiteName) == 0 {
// 		fmt.Println("empty")
// 	}

// 	parsed, err_ := json.Marshal(data.Websites)

// 	if err_ != nil {
// 		http.Error(w, "Error in parsing output: "+err_.Error(), http.StatusInternalServerError)
// 	}

// 	w.Write(parsed)
// }
