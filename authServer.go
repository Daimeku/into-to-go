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
	ConnectionString = "root:@tcp(localhost:3306)/mydb?parseTime=true"
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

	//create an auth server conf and configure it
	serverConf := osin.NewServerConfig()
	// serverConf.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	// serverConf.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE,
	// 	osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
	// serverConf.AllowGetAccessRequest = true
	// serverConf.AllowClientSecretInParams = true //ensure this is set, unless using HTTP basic auth for client secret

	//create a new osin server
	authServer := osin.NewServer(serverConf, mysqlStore)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "the home page")
	})

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

		//check for and log response errors
		if response.IsError {
			fmt.Println("there was an error - ", response.InternalError)
		}
		osin.OutputJSON(response, writer, request)
	})

	//handle the token route
	//once the authorization grant is received, send it here and get your token
	http.HandleFunc("/token", func(writer http.ResponseWriter, request *http.Request) {
		response := authServer.NewResponse()
		defer response.Close()

		//handle the access request
		accessRequest := authServer.HandleAccessRequest(response, request)
		if accessRequest != nil {
			accessRequest.Authorized = true //authorize the client
			authServer.FinishAccessRequest(response, request, accessRequest)
		}

		//check for and log response errors
		if response.IsError {
			fmt.Println("there was an error - ", response.InternalError)
		}
		//render the repsonse
		osin.OutputJSON(response, writer, request)
	})

	//start the server
	http.ListenAndServe(":8888", nil)
}
