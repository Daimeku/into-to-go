package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	// "github.com/gorilla/context"
	"net/http"
)

var sessionStore = sessions.NewCookieStore([]byte("testing"))

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe("localhost:8888", nil)
}

func handleIndex(writer http.ResponseWriter, request *http.Request) {
	session, err := sessionStore.Get(request, "test-session")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	//if the user already had a sessionId
	if session.Values["sessionId"] != nil {
		fmt.Fprintf(writer, "You had a cookie %s", session.Values["test-key"])
	} else {
		//populate the session values
		session.Values["test-key"] = "test value"
		session.Values["sessionId"] = 1
		session.Save(request, writer)
		fmt.Fprintf(writer, "You didnt have a cookie")
	}

}
