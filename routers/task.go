package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/marcusbello/gotaskmanager/common"
	"github.com/marcusbello/gotaskmanager/controllers"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	taskRouter.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	taskRouter.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.GetTaskById).Methods("GET")
	taskRouter.HandleFunc("/tasks/users/{id}", controllers.GetTaskByUser).Methods("GET")
	taskRouter.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandleFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))

	return router
}
