package main

import("net/http";
"io/ioutil";
	"fmt";
	"strings";
	"errors";
	"regexp";
)

const MAXCHATS = 15
const PAGEPATH = "/html/"
const LOGPATH = "/"


type Page struct {
	Title string;
	Body []byte;
	Path string;
	Links []string;
}

var validPath = regexp.MustCompile("^/([a-zA-Z0-9.-]+)$")

//validates proper addresses
func getTitle(res http.ResponseWriter, req *http.Request, valid *regexp.Regexp) (string, error) {
	m := valid.FindStringSubmatch(req.URL.Path)
	if m == nil {
		http.NotFound(res, req)
		return "", errors.New("Invalid Page Title")
	}
	return m[len(m) - 1], nil
}

func (p *Page) save() error {
	filename := p.Title
	//[1:] cuts out initial /
	return ioutil.WriteFile(p.Path[1:] + filename, p.Body, 0600)
}

//reads in files
func loadPage(path, title string) (*Page, error) {
	filename := title
	body, err := ioutil.ReadFile(path[1:] + filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, Path: path}, nil
}

//for regular pages, just outputs html 
func pageHandler(res http.ResponseWriter, req *http.Request){
	fmt.Printf("%s\n", req.URL.Path)
	title, err := getTitle(res, req, validPath)
	if(err != nil) { return }

	path := req.URL.Path[ : len(req.URL.Path) - len(title)]
	if(path == "/") { path = PAGEPATH } //normal pages

	page, err := loadPage(path, title)
	if(err != nil){ return }
	fmt.Fprintf(res, "%s", page.Body)
}

//outputs the logfile
func readHandler(res http.ResponseWriter, req *http.Request){
	page, err := loadPage(LOGPATH, "log")
	if(err != nil){
		fmt.Fprintf(res, "%s", "CHATLOG NOT FOUND")
		return
	}
	output := strings.Split(string(page.Body),"\n")
  if(len(output) > MAXCHATS){
    output = output[0 : MAXCHATS]
  }
	for i := len(output) - 1; i >=0; i-- {
		fmt.Fprintf(res, "%s\n", output[i])
	}
}

//Prepends message to logfile 
func writeHandler(res http.ResponseWriter, req *http.Request){
	page, err := loadPage(LOGPATH, "log")
	if(err != nil){ return }
	body := req.FormValue("say")
	page.Body = []byte(body + "\n" + string(page.Body))
	page.save()
}

func setRead(args ...string){
	for _, v := range args {
		http.Handle(v, http.FileServer(http.Dir("")))
	}
}

func main(){
	setRead("/js/", "/css/", "/img/");
	http.HandleFunc("/write", writeHandler) //handles writing
	http.HandleFunc("/read",  readHandler) //handles reading
	http.HandleFunc("/wiki/", wikiHandler) //use the wiki
	http.HandleFunc("/", pageHandler) //default
  fmt.Println(http.ListenAndServe(":80", nil))  //port 80 is http
}
