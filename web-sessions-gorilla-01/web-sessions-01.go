// continue from html-templates-01
package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// seems more clear system
func index(w http.ResponseWriter, r *http.Request) {

	// Define a struct for sending data to templates
	type TplData struct {
		InfoA string
		DataA map[int]string
		T     map[string]string
	}

	// init struct
	tplData := new(TplData)
	tplData.DataA = make(map[int]string)
	tplData.T = make(map[string]string)

	// these should be global in production for performance - template caching
	var htmlTpl = template.Must(template.ParseGlob("tpl/*.html"))

	// Add some data
	tplData.InfoA = "titleA"

	tplData.DataA[1] = "a"
	tplData.DataA[2] = "b"
	tplData.DataA[3] = "c"

	tplData.T["txt1"] = "bmm"

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	tplData.T["txt1"] = (session.Values["foo"].(string))

	// Process template and write to response to client
	err = htmlTpl.ExecuteTemplate(w, "index.html", tplData)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static-dir"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
