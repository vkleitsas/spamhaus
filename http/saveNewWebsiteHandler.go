package http

import (
	"assignment/app"
	"assignment/domain"
	"encoding/json"
	"net/http"
)

type SaveNewWebsiteUrlHandler struct {
	crudService app.CrudService
}

func NewSaveNewWebsiteUrlHandler(crudService app.CrudService) *SaveNewWebsiteUrlHandler {
	return &SaveNewWebsiteUrlHandler{
		crudService: crudService,
	}
}

type SaveNewWebsiteUrlHandlerInterface interface {
	SaveNewWebsiteUrl(w http.ResponseWriter, r *http.Request)
}

func (s *SaveNewWebsiteUrlHandler) SaveNewWebsiteUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request domain.URLEntry

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.crudService.SaveWebsiteURL(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
