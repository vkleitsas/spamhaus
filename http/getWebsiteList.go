package http

import (
	"assignment/app"
	"encoding/json"
	"net/http"
)

type GetWebsiteUrlListHandler struct {
	crudService app.CrudService
}

func NewGetWebsiteListUrlHandler(crudService app.CrudService) *GetWebsiteUrlListHandler {
	return &GetWebsiteUrlListHandler{
		crudService: crudService,
	}
}

type GetWebsiteUrlListHandlerInterface interface {
	SaveNewWebsiteUrl(w http.ResponseWriter, r *http.Request)
}

func (s *SaveNewWebsiteUrlHandler) GetWebsiteUrlList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sort := r.URL.Query().Get("sort")

	result, err := s.crudService.GetWebsiteURLS(sort)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
