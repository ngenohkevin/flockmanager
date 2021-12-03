package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	L      *log.Logger
	Db     *gorm.DB
}

func (a *App) SetRouters() http.Handler {
	//routing for users
	a.Post("/v1/kukuchic/register", a.handleRequest(a.Register))
	a.Post("/v1/kukuchic/login", a.handleRequest(a.Login))

	// Routing for handling the Kuroilers
	a.Get("/v1/kukuchic/kuroiler", a.handleRequest(a.GetAllKuroilers))
	a.Post("/v1/kukuchic/kuroiler", a.handleRequest(a.CreateKuroiler))
	a.Get("/v1/kukuchic/kuroiler/{title}", a.handleRequest(a.GetKuroiler))
	a.Put("/v1/kukuchic/kuroiler/{title}", a.handleRequest(a.UpdateKuroiler))
	a.Delete("/v1/kukuchic/kuroiler/{title}", a.handleRequest(a.DeleteKuroiler))

	// Routing for handling the RainbowRoosters
	a.Get("/v1/kukuchic/rainbowrooster", a.handleRequest(a.GetAllRainbowRoosters))
	a.Post("/v1/kukuchic/rainbowrooster", a.handleRequest(a.CreateRainbowRooster))
	a.Get("/v1/kukuchic/rainbowrooster/{title}", a.handleRequest(a.GetRainbowRooster))
	a.Put("/v1/kukuchic/rainbowrooster/{title}", a.handleRequest(a.UpdateRainbowRooster))
	a.Delete("/v1/kukuchic/rainbowrooster/{title}", a.handleRequest(a.DeleteRainbowRooster))

	// Routing for handling the Broilers
	a.Get("/v1/kukuchic/broiler", a.handleRequest(a.GetAllBroilers))
	a.Post("/v1/kukuchic/broiler", a.handleRequest(a.CreateBroiler))
	a.Get("/v1/kukuchic/broiler/{title}", a.handleRequest(a.GetBroiler))
	a.Put("/v1/kukuchic/broiler/{title}", a.handleRequest(a.UpdateBroiler))
	a.Delete("/v1/kukuchic/broiler/{title}", a.handleRequest(a.DeleteBroiler))

	// Routing for handling the Layers
	a.Get("/v1/kukuchic/layer", a.handleRequest(a.GetAllLayers))
	a.Post("/v1/kukuchic/layer", a.handleRequest(a.CreateLayer))
	a.Get("/v1/kukuchic/layer/{title}", a.handleRequest(a.GetLayer))
	a.Put("/v1/kukuchic/layer/{title}", a.handleRequest(a.UpdateLayer))
	a.Delete("/v1/kukuchic/layer/{title}", a.handleRequest(a.DeleteLayer))

	// Routing for handling Flocks
	a.Get("/v1/kukuchic/flock", a.handleRequest(a.GetAllFlocks))
	a.Get("/v1/kukuchic/flock/{id:[0-9]+}", a.handleRequest(a.GetFlock))
	//a.Post("/v1/kukuchic/flock", a.handleRequest(a.CreateFlock))

	//return a.Router
	return a.enableCORS(a.Router)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.Db, w, r)
	}
}
