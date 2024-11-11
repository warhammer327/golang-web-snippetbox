package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"warhammer327.github.io/snippetbox/pkg/forms"
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
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as the data.
	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}
	// Because the form data (with type url.Values) has been anonymously embedded
	// in the form.Form struct, we can use the Get() method to retrieve
	// the validated value for a particular form field.
	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
