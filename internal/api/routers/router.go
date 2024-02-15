package router

import (
	"github.com/alrizkipascar/go-jwt/internal/api/handlers"
	middlewares "github.com/alrizkipascar/go-jwt/internal/api/middleware"
	"github.com/gorilla/mux"
)

// type APIServer struct {
// 	listenAddr string
// 	store      Storage
// }

// func NewAPIServer(listenAddr string, store Storage) *APIServer {
// 	return &APIServer{
// 		listenAddr: listenAddr,
// 		store:      store,
// 	}
// }

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", middlewares.MakeHTTPHandleFunc(handlers.Login))
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.CreateAccount)).Methods("POST")
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.GetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", middlewares.WithJWTAuth(middlewares.MakeHTTPHandleFunc(handlers.GetAccountByID)))
	router.HandleFunc("/transfer", middlewares.MakeHTTPHandleFunc(handlers.Transfer))
	// log.Println("JSON API server runnin on port: ", handlers.listenAddr)

	return router
	// http.ListenAndServe(s.listenAddr, router)

}
