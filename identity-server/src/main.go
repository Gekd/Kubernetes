package main

import (
	"context"
	"flag"
	"fmt"
	"identity-server/src/handlers"
	"identity-server/src/variables"
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/golang-jwt/jwt"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func init() {
	flag.BoolVar(&variables.DumpRequestsAndResponses, "d", true, "Dump requests and responses")
	flag.StringVar(&variables.ClientID, "i", "222222", "The client id being passed in")
	flag.StringVar(&variables.ClientSecret, "s", "22222222", "The client secret being passed in")
	flag.StringVar(&variables.RedirectUriDomain, "r", "http://localhost:9094", "The domain of the redirect url")
	flag.IntVar(&variables.ServerPort, "p", 9096, "the base port for the server")
}

func main() {
	flag.Parse()
	if variables.DumpRequestsAndResponses {
		log.Println("Dumping requests")
	}
	authorizationManager := manage.NewDefaultManager()
	authorizationManager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	authorizationManager.MustTokenStorage(store.NewMemoryTokenStore())

	authorizationManager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set(variables.ClientID, &models.Client{
		ID:     variables.ClientID,
		Secret: variables.ClientSecret,
		Domain: variables.RedirectUriDomain,
	})
	authorizationManager.MapClientStorage(clientStore)

	authorizationServer := server.NewServer(server.NewConfig(), authorizationManager)

	authorizationServer.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})

	authorizationServer.SetUserAuthorizationHandler(handlers.UserAuthorizeHandler)

	authorizationServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	authorizationServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/auth", handlers.AuthHandler)

	http.HandleFunc("/oauth/authorize", func(responseWriter http.ResponseWriter, request *http.Request) {
		handlers.AuthorizeHandler(responseWriter, request, authorizationServer)
	})

	http.HandleFunc("/oauth/token", func(responseWriter http.ResponseWriter, request *http.Request) {
		handlers.TokenHandler(responseWriter, request, authorizationServer)
	})

	http.HandleFunc("/test", func(responseWriter http.ResponseWriter, request *http.Request) {
		handlers.IntrospectionHandler(responseWriter, request, authorizationServer)
	})

	log.Printf("Server is running at %d port.\n", variables.ServerPort)
	log.Printf("Point your OAuth client Auth endpoint to %s:%d%s", "http://localhost", variables.ServerPort, "/oauth/authorize")
	log.Printf("Point your OAuth client Token endpoint to %s:%d%s", "http://localhost", variables.ServerPort, "/oauth/token")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", variables.ServerPort), nil))
}
