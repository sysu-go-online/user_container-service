package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/user_container-service/controller"
	"github.com/sysu-go-online/public-service/types"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	// user collection
	r.Handle("/users/{username}/servers/{id}/ports/{portNumber}", types.ErrorHandler(controller.ChangePortStatusHandler)).Methods("PATCH")
	r.Handle("/users/{username}/servers/{id}/ports", types.ErrorHandler(controller.GetPortsStatusHandler)).Methods("GET")
	r.Handle("/users/{username}/servers/{id}/status", types.ErrorHandler(controller.ChangeContainerStatusHandler)).Methods("PATCH")
	r.Handle("/users/{username}/servers/{id}", types.ErrorHandler(controller.RemoveContainerHandler)).Methods("DELETE")
	r.Handle("/users/{username}/servers", types.ErrorHandler(controller.CreateContainerHandler)).Methods("POST")
	r.Handle("/users/{username}/servers", types.ErrorHandler(controller.GetAllContainersStatusHandler)).Methods("GET")

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}