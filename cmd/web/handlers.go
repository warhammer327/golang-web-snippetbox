package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
	"warhammer327.github.io/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
	// for _, snippet := range s{
	// 	fmt.Fprintf(w,"%v\n",snippet)
	// }
	//	data := &templateData{Snippets: s}
	//	files := []string{
	//		"./ui/html/home.page.tmpl",
	//		"./ui/html/base.layout.tmpl",
	//		"./ui/html/footer.partial.tmpl",
	//	}
	//	ts, err := template.ParseFiles(files...)
	//	if err != nil {
	//		app.serverError(w, err)
	//		return
	//	}
	//
	//	err = ts.Execute(w, data)
	//	if err != nil {
	//		app.serverError(w, err)
	//		return
	//	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
	}
	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
	// data := &templateData{Snippet: s}
	//
	//	files := []string{
	//		"./ui/html/show.page.tmpl",
	//		"./ui/html/base.layout.tmpl",
	//		"./ui/html/footer.partial.tmpl",
	//	}
	//
	// ts, err := template.ParseFiles(files...)
	//
	//	if err != nil {
	//		app.serverError(w, err)
	//		return
	//	}
	//
	// err = ts.Execute(w, data)
	//
	//	if err != nil {
	//		app.serverError(w, err)
	//		return
	//	}
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}
	// Check that the Content field isn't blank.
	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}
	// Check the expires field isn't blank and matches one of the permitted
	// values ("1", "7" or "365").
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}
	// If there are any errors, dump them in a plain text HTTP response and return
	// from the handler.
	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	if len(errors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{FormErrors: errors, FormData: r.Form})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
