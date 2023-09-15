package main

import (
	"net/http"
	"text/template"

	"ascii-art-web/funcs/mainfunc"
)

func Base(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		printErr(w, http.StatusNotFound)
		return
	}
	tpl, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusInternalServerError), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(500)
		// e.Execute(w, txt)
		printErr(w, http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusInternalServerError), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(500)
		// e.Execute(w, txt)
		printErr(w, http.StatusInternalServerError)
		return
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusMethodNotAllowed), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(405)
		// e.Execute(w, txt)
		printErr(w, http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	res, err := mainfunc.WriteStr(text, banner)
	if err != nil {
		// http.Error(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusNotImplemented), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(501)
		// e.Execute(w, txt)
		printErr(w, http.StatusNotImplemented)
		return
	}

	// res := result{text, err}
	tpl, err := template.ParseFiles("ui/html/ascii-art.html")
	if err != nil {
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusInternalServerError), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(500)
		// e.Execute(w, txt)
		printErr(w, http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, res)
	if err != nil {
		// txt, _ := mainfunc.WriteStr(http.StatusText(http.StatusInternalServerError), "standard")
		// e := template.Must(template.ParseFiles("./ui/html/error.html"))
		// // r.Header.Set(400, "")
		// w.WriteHeader(500)
		// e.Execute(w, txt)
		printErr(w, http.StatusInternalServerError)
		return
	}
}

func printErr(w http.ResponseWriter, StatusCode int) {
	txt, _ := mainfunc.WriteStr(http.StatusText(StatusCode), "standard")
	e := template.Must(template.ParseFiles("./ui/html/error.html"))
	// r.Header.Set(400, "")
	w.WriteHeader(StatusCode)
	e.Execute(w, txt)
}
