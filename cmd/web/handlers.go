package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/jiangzhifang/snippetbox/pkg/forms"
	"github.com/jiangzhifang/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	// panic("Oops! something went wrong")
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})

}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})

}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !Form.Valid(){
		app.render(w,r,"create.page.tmpl", &templateData{Form:form})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)

}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Create a new snippet..."))
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil)
	})
}
