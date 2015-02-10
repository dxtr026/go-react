package main

import (
	"log"
	"html/template"
	"net/http"
	"path"
	"os"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "os/exec"
	// "fmt"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", r.URL.Path)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}


	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func main() {

	// handler for static files --> css, and libs
	build := http.FileServer(http.Dir("build"))
	http.Handle("/build/", http.StripPrefix("/build/", build))

	lib := http.FileServer(http.Dir("lib"))
	http.Handle("/lib/", http.StripPrefix("/lib/", lib))

	// main routing of files
	http.HandleFunc("/", serveTemplate)

	// converting all jsx into js files , commenting out this for now as this is shifted to go-reload
	//    cmd := exec.Command("jsx","src/","build/")
	//    out, err := cmd.Output()
	//    if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(string(out))

	// starting server
	log.Println("Listening... on 8080")
	http.ListenAndServe(":8080", nil)
}