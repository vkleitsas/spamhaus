package http

import "github.com/gorilla/mux"

type Router struct {
	websiteRoutes WebsiteRoutesInterface
}

func InitMainRouter(w WebsiteRoutesInterface) *Router {
	return &Router{
		websiteRoutes: w,
	}
}

func (r *Router) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = r.websiteRoutes.SetWebsiteRoutes(router)
	return router
}
