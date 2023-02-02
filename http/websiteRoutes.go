package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebsiteRoutes struct {
	saveNewWebsiteUrlHandler SaveNewWebsiteUrlHandler
}

type WebsiteRoutesInterface interface {
	SetWebsiteRoutes(router *mux.Router) *mux.Router
}

func InitWebsiteRoutes(save SaveNewWebsiteUrlHandler) WebsiteRoutes {
	return WebsiteRoutes{
		saveNewWebsiteUrlHandler: save,
	}
}
func (w WebsiteRoutes) SetWebsiteRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api", w.saveNewWebsiteUrlHandler.SaveNewWebsiteUrl).Methods(http.MethodPost)
	router.HandleFunc("/api", w.saveNewWebsiteUrlHandler.GetWebsiteUrlList).Methods(http.MethodGet)
	return router
}
