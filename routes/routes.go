package routes

import(
	"github.com/Sally90/Embedding2/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router){
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/blog", controllers.Blog).Methods("GET")
}
