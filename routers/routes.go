package routes

import(
	"github.com/gorilla/mux"
	"github.com/your/uraj323/bookstore_apis_mongo/pkg/api"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	api.SetupBookRoutes(router)

	return router
}