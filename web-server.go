package main

import (
	 "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

var pageTemplates *template.Template

func main() {
	page := Page{ "Home", []byte("The home page")}
	page.Save()
	pageFromFile, _ := loadPage("Home")
	fmt.Println(string(pageFromFile.Body))

	parseTemplates()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/error/", errorHandler)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title string
	Body []byte
}

//saves a pages contents to a file
func (page *Page) Save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)	//write to the file with 0600 permission (for this user only)
}

//creates a page and reads the Body from file
func loadPage(name string) (*Page, error) {
	filename := name + ".txt"	// create the filename
	var page *Page = &Page{} 	//create the page and error variables
	var err error
	(*page).Title = name	

	page.Body, err = ioutil.ReadFile(filename)	//read the body from file
	if err != nil {								//check for errors
		fmt.Println("there was an error reading the file - ", err)
		return nil, err
	}

	return page, nil
}

//handles all routes with /view/- loads the data from the appropriate txt file and displays it as HTML
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):] // the title is the remainder of the URL after /view/
	page, err := loadPage(title)
	if err != nil {
		http.Error(writer, "The requested page does not exist", http.StatusInternalServerError)
		return
	}
	renderTemplate(writer, "viewPageTemplate", page)
}

//handles all routes with /edit/- returns the edit template with the requested page details 
func editHandler(writer http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/edit/"):] 
	page, err := loadPage(title)
	if err != nil {
		redirectToError(writer, request)
		fmt.Println("There was an error loading the page for editing - ", err)
		return
	}
	renderTemplate(writer, "editFormTemplate", page)

}

//handles POST to save pages
func saveHandler(writer http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(writer, (request.Form["title"]))
	title := request.URL.Path[len("/save/"):]
	err := request.ParseForm()
	if err != nil {
		redirectToError(writer, request)
		return
	}
	fmt.Println(request.Form)
	page, err := loadPage(title)
	if err!=nil {
		redirectToError(writer, request)
		return
	}
	page.Title = strings.Join(request.Form["title"],"")
	body := strings.Join(request.Form["body"],"")
	page.Body = []byte(body)
	err= page.Save()
	if err!= nil {
		fmt.Println("Page not saved")
		return
	}
	fmt.Println("file saved - ", page.Title)

	//delete old file
}

//displays an error message for the user - handles all routes with /error/
func errorHandler(writer http.ResponseWriter, request *http.Request){
	//displays an error message for the user
	fmt.Fprintf(writer, "There has been an error")
}

//redirects the user to the error page
func redirectToError(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/error/", http.StatusSeeOther)
}

//parses and executes the given page template
func renderTemplate(writer http.ResponseWriter, templateName string, page *Page) {
	// pageTemplate,err := template.ParseFiles(templateName+".html")
	err := pageTemplates.ExecuteTemplate(writer, templateName+".html", page)
	if err != nil {
		fmt.Println("There was an error loading the page", err)
		return
	}
	// pageTemplate.Execute(writer, page)
}

//parse all templates
func parseTemplates()  {
	pageTemplates = template.Must(template.ParseFiles("editFormTemplate.html", "viewPageTemplate.html", "errorPageTemplate.html"))
}