package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		// r is set to 400 inside of function
		app.notFound(w)
		return
	}
	// files := []string{
	// 	"./ui/html/home.page.html",
	// 	"./ui/html/base.layout.html",
	// 	"./ui/html/footer.partial.html",
	// }

	files := []string{"./ui/html/home.index.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}
}

func (app *application) showBlogs(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/blogs" {
	// 	// http.NotFound(w, r)
	// 	// r is set to 400 inside of function
	// 	app.notFound(w)
	// 	return
	// }

	files := []string{"./ui/html/blog.landing.html"}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}
}

func (app *application) showPosts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	postCache := []int{1, 2, 3}

	_, ok := findInt(postCache, id)
	if !ok {
		app.notFound(w)
		return
	}

	post := fmt.Sprintf("./ui/html/blog.post%d.html", id)
	files := []string{post}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		// app.errorLog.Println(err.Error())
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err)
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	w.Write([]byte("Create a new snippet..."))
// }
