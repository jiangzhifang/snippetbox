package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	// "github.com/gobuffalo/packr/v2"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanaic, app.logRequest, secureHeaders)

	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	// mux := http.NewServeMux()
	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	// fsbox := packr.New("someBoxName", "./ui/static")
	fileServer := http.FileServer(http.Dir("./ui/static"))
	// fileServer := http.FileServer(fsbox)
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	// mux.Get("/static/", fileServer)

	// return app.recoverPanaic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(mux)
}
