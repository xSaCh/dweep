package pkg

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xSaCh/dweep/pkg/handlers/watchlist"
	"github.com/xSaCh/dweep/pkg/storage"
)

type APIServer struct {
	addr    string
	storage storage.Storage
}

func NewAPIServer(addr string, storage storage.Storage) *APIServer {
	return &APIServer{
		addr:    addr,
		storage: storage,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// productStore := product.NewStore(s.db)
	// productHandler := product.NewHandler(productStore, userStore)
	// productHandler.RegisterRoutes(subrouter)
	watchlistSubRouter := router.PathPrefix("/watchlist/").Subrouter()
	watchlist.NewHandler(s.storage).RegisterRoutes(watchlistSubRouter)

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(`{"message": "okie"}`))
	})

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
