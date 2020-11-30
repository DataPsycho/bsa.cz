package main

import "net/http"

func (app *application) routers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/blogs/", app.showBlogs)
	mux.HandleFunc("/blogs/posts/", app.showPosts)
	// mux.HandleFunc("snippet", app.showSnippet)
	// mux.HandleFunc("/snippet/create", app.createSnippet)
	// mux.HandleFunc("", )

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
