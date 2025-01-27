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

//===========================================================================================================================================================================================================================================================================================================================================================================

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type requestData struct {
	Websites []string `json:"websites"`
}

type StatusChecker interface {
	Check(ctx context.Context, name string) (status bool, err error)
	CheckMap(ctx context.Context, wg *sync.WaitGroup, m *sync.Mutex, name string, siteStatus map[string]string)
}

type httpChecker struct{}

type websitesHandler struct {
	ctx           context.Context
	sitesStatus   map[string]string
	statusChecker StatusChecker
	m             sync.Mutex
	wg            sync.WaitGroup
}

func (h *httpChecker) Check(ctx context.Context, name string) (status bool, err error) {
	resp, err := http.Get(name)
	if err != nil {
		return false, errors.New(err.Error())
	}
	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, errors.New("status code doesn't match")
}

func (h *httpChecker) CheckMap(ctx context.Context, wg *sync.WaitGroup, m *sync.Mutex, name string, siteStatus map[string]string) {
	defer wg.Done()
	m.Lock()
	resp, err := http.Get(name)

	if err != nil {
		siteStatus[name] = "DOWN"
		m.Unlock()
		return
	}
	if resp.StatusCode == 200 {
		siteStatus[name] = "UP"
		m.Unlock()
		return
	}
	siteStatus[name] = "DOWN"
	m.Unlock()
}

func (h *websitesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	// if GET method
	if method == http.MethodGet {

		websiteName := r.URL.Query().Get("name")

		h.handleGetRequest(websiteName, w)

	} else {
		// if POST request

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Problem reading request body: "+err.Error(), http.StatusInternalServerError)
		}

		var resp requestData
		newErr := json.Unmarshal(body, &resp)

		if newErr != nil {
			http.Error(w, "Found Error: "+newErr.Error(), http.StatusInternalServerError)
		}

		// constantly prints Status in terminal at an interval of 10 seconds..
		for {
			for _, site := range resp.Websites {
				h.wg.Add(1)
				go h.statusChecker.CheckMap(h.ctx, &h.wg, &h.m, site, h.sitesStatus)
			}
			h.wg.Wait()
			fmt.Println(h.sitesStatus)
			returnHttpResponse(w, h.sitesStatus)
			time.Sleep(time.Second * 10)
		}
	}
}

func (handler *websitesHandler) handleGetRequest(websiteName string, w http.ResponseWriter) {

	// if `name` param doesn't exists, check for all sites..
	if len(websiteName) == 0 {
		sitesString := []string{
			"https://www.facebook.com",
			"https://www.google.com",
			"https://www.fakewebsite1.com"}

		for _, site := range sitesString {
			handler.wg.Add(1)
			go handler.statusChecker.CheckMap(handler.ctx, &handler.wg, &handler.m, site, handler.sitesStatus)
		}

		handler.wg.Wait()
		returnHttpResponse(w, handler.sitesStatus)
		return
	}

	// if name param exists
	handler.wg.Add(1)
	go handler.statusChecker.CheckMap(handler.ctx, &handler.wg, &handler.m, websiteName, handler.sitesStatus)
	handler.wg.Wait()
	returnHttpResponse(w, handler.sitesStatus)
}

func returnHttpResponse(w http.ResponseWriter, data any) {
	response, marshalErr := json.Marshal(data)
	if marshalErr != nil {
		http.Error(w, "Error: "+marshalErr.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func main() {

	mux := http.DefaultServeMux
	ctx := context.Background()
	var m sync.Mutex
	var wg sync.WaitGroup
	httpChecker := httpChecker{}

	sitesStatus := make(map[string]string)

	ctx = context.WithValue(ctx, "SiteStatus", sitesStatus)
	handler := websitesHandler{
		ctx,
		sitesStatus,
		&httpChecker,
		m,
		wg,
	}

	mux.Handle("GET /websites", &handler)
	mux.Handle("POST /websites", &handler)

	fmt.Println("Listen to Port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
