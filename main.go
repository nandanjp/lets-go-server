package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...\n", id)
	// w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost) //show allowed headers; must be called before the other two methods
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusNotAcceptable) //essentially wrapper calling the two above methods for us
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name": "Alex"}`))
	// w.Write([]byte("Create a new snippet...."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)                    //subtree path: basically acts like /**, acts like catch-all
	mux.HandleFunc("/snippet/view", snippetView) //fixed path: doesn't end in /, meaning this route must be matched exactly
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Print("Error in starting the server for some reason......")
		os.Exit(1)
	}

	// log.Fatal(err)
}
