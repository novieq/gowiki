
package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
)

//Body element is a byte rather than a string because that is what io libraries will expect
type Page struct {
    Title string
    Body  []byte
}
//save method returns error value because that is the return type of WriteFile
//0600 means that the file should be created with read write permissions for the current user only
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile( filename, p.Body,0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page { Title:title, Body:body}, err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p,err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	//the viewHandler will get registered against this pattern in the DefaultServerMux
	http.HandleFunc("/view/",viewHandler)
	http.HandleFunc("/edit/",editHandler)
	http.HandleFunc("/save/",saveHandler)
	http.ListenAndServe(":8080",nil)
}
/*
func main() {
	p1 := &Page{Title:"NewPage",Body:[]byte("This is my first page")}
	_ = p1.save()
	p2,_ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}*/