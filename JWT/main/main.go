package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type blog struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allBlogs []blog

var blogs = allBlogs{
	{
		ID:      1,
		Name:    "Blog 1",
		Content: "Esto es un blog de prueba",
	},
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blogs)
}

func createBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog blog
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Blog")
	}

	json.Unmarshal(reqBody, &newBlog)

	newBlog.ID = len(blogs) + 1
	blogs = append(blogs, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBlog)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my ApiBlog")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/blogs", getBlogs)

	log.Fatal(http.ListenAndServe(":3000", router))
}
