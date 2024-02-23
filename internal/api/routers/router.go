package router

import (
	"net/http"

	"github.com/alrizkipascar/go-jwt/internal/api/handlers"
	middlewares "github.com/alrizkipascar/go-jwt/internal/api/middleware"
	"github.com/gorilla/mux"
)

// type APIServer struct {
// 	listenAddr string
// 	store      Storage
// }

//	func NewAPIServer(listenAddr string, store Storage) *APIServer {
//		return &APIServer{
//			listenAddr: listenAddr,
//			store:      store,
//		}
//	}
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("MIDDLEWARE CALLED")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-jwt-token, user-id")
		next.ServeHTTP(w, r)
	})
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/login", middlewares.MakeHTTPHandleFunc(handlers.Login))
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.CreateAccount)).Methods("POST")
	router.HandleFunc("/account", middlewares.MakeHTTPHandleFunc(handlers.GetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", middlewares.WithJWTAuth(middlewares.MakeHTTPHandleFunc(handlers.GetAccountByID)))
	router.HandleFunc("/transfer", middlewares.MakeHTTPHandleFunc(handlers.Transfer))
	// log.Println("JSON API server runnin on port: ", handlers.listenAddr)

	return router
	// http.ListenAndServe(s.listenAddr, router)

}
