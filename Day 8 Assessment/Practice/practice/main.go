package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ====================================================================================================================================================

// Training Session Practice code

// type Post struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name,omitempty"`
// }

// func main() {
// 	mux := http.DefaultServeMux

// 	mux.HandleFunc("GET /posts", listPostsHandler)
// 	http.HandleFunc("POST /posts/{id}", getPostHandler)
// 	fmt.Println("Server is running at http://localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }

// func getPostHandler(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading request body", http.StatusBadRequest)
// 		return
// 	}

// 	var post Post
// 	err = json.Unmarshal(body, &post)
// 	if err != nil {
// 		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
// 		return
// 	}
// 	post.ID = 1
// 	handleResponse(w, post, http.StatusOK)
// }

// func listPostsHandler(w http.ResponseWriter, r *http.Request) {

// 	posts := []Post{
// 		{ID: 1, Name: "First post"},
// 		{ID: 2, Name: "Second post"},
// 	}

// 	handleResponse(w, posts, http.StatusOK)
// }

// func handleResponse(w http.ResponseWriter, message any, statusCode int) {
// 	response, err := json.Marshal(message)
// 	if err != nil {
// 		http.Error(w, "Error converting response to json", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(statusCode)
// 	w.Write(response)
// }

// ========================================================================================================================================================================

// Use of HANDLE + HANDLE_FUNC and HANDLER + HANDLER_FUNC (http package)

// code here helps in understand difference between the use of Handler + HandlerFunc and Handle and HandleFunc
// handle is basically a struct type and takes in two params where first one is the pattern and other is HandlerFunc, which is any function that takes in http.responseWriter and http.Request
// whereas handler is an interface and has only one methods which is ServeHTTP, hence any struct that follow this function can be passed as an handleFunction.

// type student struct {
// 	name string
// 	year int
// 	prn  int
// }

// func (st *student) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	st1 := student{"Leo", 10, 10}

// 	parsed, _ := json.Marshal(st1)

// 	w.Write(parsed)
// }

// func main() {

// 	mux := http.DefaultServeMux

// 	st1 := student{"Hehe", 10, 109}

// 	// http.Handle("/", http.)
// 	mux.Handle("/class", &st1)                          // here ignore &st1, here st1 is basically just a HANDLER (interface with 1 methods - ServeHTTP)
// 	http.HandleFunc("GET /students", getRequestHandler) // here getRequestHandler is a Handler_Func
// 	fmt.Println("Listening to Port 8080")
// 	log.Fatal(http.ListenAndServe(":3000", mux))
// }

// func getRequestHandler(w http.ResponseWriter, r *http.Request) {
// 	// students := []student{{"Harsh", 3, 2014}, {"Leo", 4, 1004}, {"haha", 10, 10}}

// 	// response, err := json.Marshal(students)

// 	// if err != nil {
// 	// 	http.Error(w, "Error in parsing to JSON", http.StatusInternalServerError)
// 	// }
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	w.Write([]byte("Hello"))
// }

//======================================================================================================================================================

func main() {
	name := "https://www.facebook.com"
	resp, err := http.Get(name)

	if err != nil {
		fmt.Println(errors.New("Error found: " + err.Error()))
	}
	fmt.Println(resp.Status)
}

// =======================================================================================================================================================

// type requestData struct {
// 	Websites []string `json:"websites"`
// }

// type StatusChecker interface {
// 	Check(ctx context.Context, name string) (status bool, err error)
// }

// type httpChecker struct {
// 	ctx         context.Context
// 	sitesStatus map[string]bool
// }

// func (h httpChecker) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	// we find what kind of request it is here... so basically if it is a get request then we StatusCheck for the website passed in param
// 	// else if its a post request that means that a list of sites are passed, hence we check for all of them using a different function maybe.. for Granularity or Modularity

// 	method := r.Method

// 	if method == "GET" {
// 		websiteName := r.URL.Query().Get("name")

// 		isUp, err := h.CheckOne(h.ctx, websiteName)

// 		if err != nil || !isUp {
// 			w.Write([]byte("DOWN"))
// 			return
// 		}
// 		w.Write([]byte("UP"))

// 	} else {

// 		body, err := io.ReadAll(r.Body)

// 		if err != nil {
// 			http.Error(w, "Problem reading request body: "+err.Error(), http.StatusInternalServerError)
// 		}

// 		var resp requestData

// 		newErr := json.Unmarshal(body, &resp)

// 		if newErr != nil {
// 			http.Error(w, "Found Error: "+newErr.Error(), http.StatusInternalServerError)
// 		}

// 		h.checkStatusAll(h.ctx, resp.Websites)

// 		response, marshalErr := json.Marshal(h.sitesStatus)
// 		if marshalErr != nil {
// 			http.Error(w, "Error: "+marshalErr.Error(), http.StatusInternalServerError)
// 		}

// 		w.Write(response)

// 	}
// }

// func (h httpChecker) checkStatusAll(ctx context.Context, siteNames []string) {
// 	for _, site := range siteNames {
// 		status, err := h.CheckOne(ctx, site)
// 		if err != nil {
// 			h.sitesStatus[site] = false
// 		}
// 		if status {
// 			h.sitesStatus[site] = true
// 		}
// 	}
// }

// func (h httpChecker) CheckOne(ctx context.Context, name string) (status bool, err error) {
// 	resp, err := http.Get(name)
// 	if err != nil {
// 		return false, errors.New(err.Error())
// 	}

// 	if resp.StatusCode == 200 {
// 		return true, nil
// 	}

// 	return false, errors.New("status code doesn't match")
// }

// func main() {

// 	mux := http.DefaultServeMux
// 	ctx := context.Background()

// 	sitesStatus := make(map[string]bool)

// 	httpObj := httpChecker{
// 		ctx,
// 		sitesStatus,
// 	}

// 	// mux.HandleFunc("GET /helloServer", handleGetRequest) // just for testing
// 	// mux.HandleFunc("POST /websites", websiteStatusHandler)

// 	mux.Handle("GET /websites", httpObj)
// 	mux.Handle("POST /websites", httpObj)

// 	fmt.Println("Listen to Port :8080")
// 	log.Fatal(http.ListenAndServe(":8080", mux))
// }

// // func handleGetRequest(w http.ResponseWriter, r *http.Request) {
// // 	_, err := w.Write([]byte("Hello"))

// // 	if err != nil {
// // 		http.Error(w, "Error in write: "+err.Error(), http.StatusInternalServerError)
// // 	}
// // }

// // func websiteStatusHandler(w http.ResponseWriter, r *http.Request) {

// // 	body, err := io.ReadAll(r.Body)

// // 	if err != nil {
// // 		http.Error(w, "Problem reading request body: "+err.Error(), http.StatusInternalServerError)
// // 	}

// // 	var data requestData

// // 	errr := json.Unmarshal(body, &data)

// // 	if errr != nil {
// // 		http.Error(w, "Error found: "+errr.Error(), http.StatusInternalServerError)
// // 	}

// // 	websiteName := r.URL.Query().Get("name")

// // 	// if no params passed check for all sites

// // 	if len(websiteName) == 0 {
// // 		fmt.Println("empty")
// // 	}

// // 	parsed, err_ := json.Marshal(data.Websites)

// // 	if err_ != nil {
// // 		http.Error(w, "Error in parsing output: "+err_.Error(), http.StatusInternalServerError)
// // 	}

// // 	w.Write(parsed)
// // }
