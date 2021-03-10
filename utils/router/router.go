package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wisnuanggoro/pokedex-web-go/config"
	"github.com/wisnuanggoro/pokedex-web-go/handlers"
	"github.com/wisnuanggoro/pokedex-web-go/utils/render"
)

type router struct {
	errorHandler  handlers.ErrorHandler
	homeHandler   handlers.HomeHandler
	detailHandler handlers.DetailHandler
	searchHandler handlers.SearchHandler
}

type Router interface {
	InitRouter(render render.Render, cfg *config.Config) *mux.Router
}

func NewRouter(
	errorHandler handlers.ErrorHandler,
	homeHandler handlers.HomeHandler,
	detailHandler handlers.DetailHandler,
	searchHandler handlers.SearchHandler) Router {
	return &router{
		errorHandler:  errorHandler,
		homeHandler:   homeHandler,
		detailHandler: detailHandler,
		searchHandler: searchHandler,
	}
}

func (rtr *router) InitRouter(render render.Render, cfg *config.Config) *mux.Router {

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", rtr.homeHandler.HomePage).Methods("GET")
	r.HandleFunc("/detail/{id}", rtr.detailHandler.DetailPage).Methods("GET")
	r.HandleFunc("/search", rtr.searchHandler.SearchPage).Methods("GET")

	// Initialize static folder
	fsImage := http.FileServer(http.Dir("./views/assets/img"))
	r.PathPrefix("/assets/img/").Handler(http.StripPrefix("/assets/img/", fsImage))
	fsCSS := http.FileServer(http.Dir("./views/assets/css"))
	r.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", fsCSS))
	fsJs := http.FileServer(http.Dir("./views/assets/js"))
	r.PathPrefix("/assets/js/").Handler(http.StripPrefix("/assets/js/", fsJs))

	return r
}
