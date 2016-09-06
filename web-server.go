package main

import (
	 "fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	page := Page{ "Home", []byte("The home page")}
	page.Save()
	pageFromFile, _ := loadPage("Home")
	fmt.Println(string(pageFromFile.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
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
		//fmt.Println("There was an error loading the page")
		return
	}
	fmt.Fprintf(writer, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
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
	pageTemplate,err := template.ParseFiles("editFormTemplate.html")
	if err != nil {
		fmt.Println("There was an error loading the edit template - ", err)
		return
	}
	pageTemplate.Execute(writer, page)

}

//displays an error message for the user
func errorHandler(writer http.ResponseWriter, request *http.Request){
	//displays an error message for the user
	fmt.Fprintf(writer, "There has been an error")
}

func redirectToError(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/error/", http.StatusSeeOther)
}