package main

import (
	"html/template"
	"net/http"
)

// seems more clear system
func indexa(w http.ResponseWriter, r *http.Request) {

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
	//tplData := &TplData{
	//	"",
	//	make(map[int]string),
	//	make(map[string]string),
	//}

	//TODO
	// maybe implement interfaces as data to be sent to templates
	//map[interface{}]interface{}

	// these should be global in production for performance - tpl caching
	//var templates = template.Must(template.ParseFiles("tpl/tpl.html"))
	var htmlTpl = template.Must(template.ParseGlob("tpl_a/*.html"))

	//log.Println(htmlTpl)

	// Add some data
	tplData.InfoA = "titleA"

	tplData.DataA[1] = "a"
	tplData.DataA[2] = "b"
	tplData.DataA[3] = "c"

	tplData.T["txt1"] = "bmm"

	// Process template and write to response to client
	err := htmlTpl.ExecuteTemplate(w, "index.html", tplData)
	if err != nil {
		//in prod replace err.error() with something else
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexb(w http.ResponseWriter, r *http.Request) {

	// these should be global in production for performance - tpl caching
	//var templates = template.Must(template.ParseFiles("tpl/tpl.html"))
	var htmlTmpl = template.Must(template.ParseGlob("tpl_b/*.html"))

	//log.Println(htmlTpl)

	m := make(map[int]string)

	m[1] = "a"
	m[2] = "b"
	m[3] = "c"

	err := htmlTmpl.ExecuteTemplate(w, "index", m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//t, err := template.ParseFiles("templates/index.html")
	//if err != nil {
	//	log.Fatal("WTF dude, error parsing your template.")
	//}

	// Render the template, writing to `w`.
	//t.Execute(w, r.Host)
}

func main() {

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static-dir"))))
	http.HandleFunc("/tpla", indexa)
	http.HandleFunc("/tplb", indexb)
	http.ListenAndServe(":8080", nil)
}
