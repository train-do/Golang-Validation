package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/service"
	"github.com/train-do/Golang-Generics/utils"
)

type HandlerDestination struct {
	Service *service.ServiceDestination
}

func NewRepoDestination(service *service.ServiceDestination) *HandlerDestination {
	return &HandlerDestination{service}
}
func (h *HandlerDestination) GetAll(w http.ResponseWriter, r *http.Request) {
	var qp model.QueryParams
	qp.Page = utils.ToInt(r.URL.Query().Get("page"))
	qp.SortDate = utils.ToBool(r.URL.Query().Get("sort_date"))
	qp.SortPrice = r.URL.Query().Get("sort_price")
	qp.SortName = utils.ToBool(r.URL.Query().Get("sort_name"))
	qp.SearchPlace = r.URL.Query().Get("search_place")
	qp.SearchDate = r.URL.Query().Get("search_date")
	qp.SearchPrice = utils.ToInt(r.URL.Query().Get("search_price"))
	// fmt.Printf("%+v\n", qp)
	data, err := h.Service.GetAll(qp)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, data, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
func (h *HandlerDestination) GetById(w http.ResponseWriter, r *http.Request) {
	id := utils.ToInt(chi.URLParam(r, "id"))
	var destination model.Destination
	err := h.Service.GetById(&destination, id)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: destination}, http.StatusOK, "")
	json.NewEncoder(w).Encode(response)
}
