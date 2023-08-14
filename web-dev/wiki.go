// simple web app
// https://go.dev/doc/articles/wiki/

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles(
	"templates/edit.html",
	"templates/view.html",
))

func (p *Page) save() error {
	filename := "./pages/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func (p *Page) load() (*Page, error) {
	filename := "./pages/" + p.Title + ".txt"
	body, err := os.ReadFile(filename)
	p.Body = body
	return p, err
}

func renderTemplate(w http.ResponseWriter, n string, p *Page) {

	e := templates.ExecuteTemplate(w, n+".html", p)

	if e != nil {
		http.Error(
			w,
			e.Error(),
			http.StatusInternalServerError,
		)
	}
}

func validate(p string) bool {
	vp := regexp.MustCompile("^/(edit|save|view)/[aA-zZ]+$")
	m := vp.FindStringSubmatch(p)
	return m != nil
}

func saveHandler(w http.ResponseWriter, r *http.Request) {

	if !validate(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")

	p := &Page{
		Title: title,
		Body:  []byte(body),
	}

	if err := p.save(); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

func editHandler(w http.ResponseWriter, r *http.Request) {

	if !validate(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	title := r.URL.Path[len("/edit/"):]

	p := &Page{
		Title: title,
		Body:  nil,
	}
	renderTemplate(w, "edit", p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	if !validate(r.URL.Path) {
		http.NotFound(w, r)
		return
	}

	title := r.URL.Path[len("/view/"):]

	p := &Page{
		Title: title,
		Body:  nil,
	}

	if p, err := p.load(); err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	} else {
		renderTemplate(w, "view", p)
	}
}

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
