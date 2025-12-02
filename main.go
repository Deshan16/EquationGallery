package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("./html"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/file1.html")
	})
	
	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/page/"):]
		
		filePath := filepath.Join("html", name)
		
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("<h1>Page not found</h1>"))
			return
		}
		
		http.ServeFile(w, r, filePath)
	})
	
	fmt.Println("Server running on http://localhost:5600")
	http.ListenAndServe(":5600", nil)
}