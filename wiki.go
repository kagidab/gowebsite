package main

import (
	"net/http";
	"html/template";
//	"fmt";
	"regexp";
	"io/ioutil";
	"math/rand";
  "strings";
	"time"
)

//paths for various functions
const WIKIPATH = "/wiki/"
const EDITPATH = WIKIPATH + "edit/"
const SAVEPATH = WIKIPATH + "save/"
const VIEWPATH = WIKIPATH + ""
const TPLPATH = "templates/"

var wikiPaths = regexp.MustCompile( "^/wiki/(save/|edit/|)([a-zA-Z0-9_ ]*)$")

//for caching templates
var templates = template.Must(template.ParseFiles(TPLPATH + "edit.tpl", TPLPATH + "view.tpl", TPLPATH + "viewall.tpl"))

// Wiki titles use spaces, underscores for paths.
// add_remove set to true makes " "->"_" and vice versa
func toggleUnderscores(str *string, add_remove bool){
  if(add_remove){
    *str = strings.Replace(*str, " ", "_", -1)
  } else {
    *str = strings.Replace(*str, "_", " ", -1)
  }
}


//goes to a wiki page, edits if unexistant
func viewHandler(res http.ResponseWriter, req *http.Request, title string) {
  toggleUnderscores(&title, true)
	page, err := loadPage(VIEWPATH, title)
	if(err != nil){ //nosuchpage, make new
		http.Redirect(res, req, EDITPATH + title, http.StatusFound)
		return
	}
	renderTemplate(res, "view", page)
}

//crates new wiki pages
func saveHandler(res http.ResponseWriter, req *http.Request, title string) {
	body := req.FormValue("body")
	page := &Page{ Title : title,  Body : []byte(body), Path : VIEWPATH }
  toggleUnderscores(&page.Title, true)
	err := page.save()
	if(err != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, VIEWPATH + title, http.StatusFound)
}

func renderTemplate(res http.ResponseWriter, tmpl string, page *Page){
  toggleUnderscores(&page.Title, false);
	err := templates.ExecuteTemplate(res, tmpl + ".tpl",
	struct{Body []byte; Links []string; Title string; SAVEPATH string; EDITPATH string;
	WIKIPATH string; VIEWPATH string;
		}{page.Body, page.Links, page.Title, SAVEPATH, EDITPATH, WIKIPATH, VIEWPATH})
	if(err != nil){
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

//edit wiki page
func editHandler(res http.ResponseWriter, req *http.Request, title string) {
  toggleUnderscores(&title, true)
	page, err := loadPage(VIEWPATH, title)
	if(err != nil){ page = &Page{ Title : title }}
	renderTemplate(res, "edit", page)
}

//outputs all the wiki pages
func listAll(res http.ResponseWriter, req *http.Request) {
	files, _ := ioutil.ReadDir("wiki/")
	body := make([]string, len(files))
	for index, f := range files {
		body[index] = f.Name()
	}
	page := Page{ Links : body, Title: "Article listing", Path: VIEWPATH}
	renderTemplate(res, "viewall", &page)
}

//selects a random page
func randompage(res http.ResponseWriter, req *http.Request) {
	files, _ := ioutil.ReadDir("wiki/");
	rand.Seed(time.Now().UTC().UnixNano())
	if(len(files) > 0){
		r := files[rand.Intn(len(files))]
		viewHandler(res, req, r.Name())
	}
}

//main handler for wiki, redirects to correct page
func wikiHandler(res http.ResponseWriter, req *http.Request) {
	title, err := getTitle(res, req, wikiPaths) //validates as well
	if(err != nil){ return }
	path := req.URL.Path[ : len(req.URL.Path) - len(title)]
	println(path)

	if(title == "List_all"){
		listAll(res, req);
	} else if(title == "Random_page"){
		randompage(res, req);
	} else if(title == ""){
    viewHandler(res, req, "Main_page")
	} else {
		switch(path){
		case EDITPATH:
			editHandler(res, req, title)
		case SAVEPATH:
			saveHandler(res, req, title)
		default:
			viewHandler(res, req, title)
		}
	}
}
