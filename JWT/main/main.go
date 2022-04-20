package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Types
type blog struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allBlogs []blog

// Persistence
var blogs = allBlogs{
	{
		ID:      1,
		Name:    "Blog One",
		Content: "Some Content",
	},
	{
		ID:      2,
		Name:    "Blog Two",
		Content: "This is the second Blog",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome the my GO API!")
}

func createBlog(w http.ResponseWriter, r *http.Request) {
	var newBlog blog
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Blog Data")
	}

	json.Unmarshal(reqBody, &newBlog)
	newBlog.ID = len(blogs) + 1
	blogs = append(blogs, newBlog)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBlog)

}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func getOneBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, blog := range blogs {
		if blog.ID == blogID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(blog)
		}
	}
}

func updateBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID, err := strconv.Atoi(vars["id"])
	var updatedBlog blog

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedBlog)

	for i, t := range blogs {
		if t.ID == blogID {
			blogs = append(blogs[:i], blogs[i+1:]...)

			updatedBlog.ID = t.ID
			blogs = append(blogs, updatedBlog)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedblog)
			fmt.Fprintf(w, "The Blog with ID %v has been updated successfully", blogID)
		}
	}

}

func deleteBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blogID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range blogs {
		if t.ID == blogID {
			blogs = append(blogs[:i], blogs[i+1:]...)
			fmt.Fprintf(w, "The blog with ID %v has been remove successfully", blogID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/blogs", createBlog).Methods("POST")
	router.HandleFunc("/blogs", getBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id}", getOneBlog).Methods("GET")
	router.HandleFunc("/blogs/{id}", deleteBlog).Methods("DELETE")
	router.HandleFunc("/blogs/{id}", updateBlog).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
