
package main

import(
	"fmt"
	"net/http"
	"strconv"
)

func home (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("hello from snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <1 {
		http.NotFound(w,r)
		return
	}
	fmt.Fprintf(w, "snippet id %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		//both code do the same
		// w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Post method mising"))
		w.Header()["Date"] = nil
		http.Error(w,"method not allowed",405)
		return
	}
	w.Write([]byte("create a specific snippet"))
}
