package routers

import (
	"github.com/gorilla/mux"
	"github.com/marcusbello/gotaskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	return router
}
