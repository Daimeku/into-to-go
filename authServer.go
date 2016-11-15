package main

import (
	"database/sql"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/felipeweb/osin-mysql"
	"net/http"
)

const (
	Driver           = "mysql"
	ConnectionString = "root:@tcp(localhost:3306)/mydb"
)

func main() {

	db, err := sql.Open(Driver, ConnectionString)
	if err != nil {
		fmt.Println("Error opening the connection - ", err)
		return
	}

	//create a new mysql store and create the schema
	mysqlStore := mysql.New(db, "osin_")
	mysqlStore.CreateSchemas()

	//create a new osin server
	authServer := osin.NewServer(osin.NewServerConfig(), mysqlStore)

	//handle the authorize route
	http.HandleFunc("/authorize", func(writer http.ResponseWriter, request *http.Request) {

		//create a response and ensure it closes at the end of this function
		response := authServer.NewResponse()
		defer response.Close()

		authRequest := authServer.HandleAuthorizeRequest(response, request)
		if authRequest != nil {

			authRequest.Authorized = true
			authServer.FinishAuthorizeRequest(response, request, authRequest)
		}
		osin.OutputJSON(response, writer, request)
	})

	http.ListenAndServe(":8888", nil)
}
