package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/akrylysov/algnhsa"
	"github.com/julienschmidt/httprouter"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	f, _ := strconv.Atoi(r.FormValue("first"))
	s, _ := strconv.Atoi(r.FormValue("second"))
	w.Header().Set("X-Hi", "foo")
	fmt.Fprintf(w, "%d", f+s)
}

func contextHandler(w http.ResponseWriter, r *http.Request) {
	proxyReq, ok := algnhsa.ProxyRequestFromContext(r.Context())
	if ok {
		resp := fmt.Sprintf("your request id: %s", proxyReq.RequestContext.RequestID)
		fmt.Fprint(w, resp)
	}
}

func main() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/index", indexHandler)
	router.HandlerFunc("GET", "/add", addHandler)
	router.HandlerFunc("GET", "/context", contextHandler)
	algnhsa.ListenAndServe(router, nil)
}

