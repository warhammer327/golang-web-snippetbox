package main

import(
	"log"
	"net/http"
)

func home (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("hello from snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("show a specific snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		// w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Post method mising"))
		w.Header()["Date"] = nil
		http.Error(w,"method not allowed",405)
		return
	}
	w.Write([]byte("create a specific snippet"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/showSnippet", showSnippet)
	mux.HandleFunc("/create/snippet", createSnippet)

	log.Println("starting server on 4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}