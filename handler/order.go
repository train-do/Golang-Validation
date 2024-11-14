package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/service"
	"github.com/train-do/Golang-Generics/utils"
)

type HandlerOrder struct {
	Service *service.ServiceOrder
}

func NewRepoOrder(service *service.ServiceOrder) *HandlerOrder {
	return &HandlerOrder{service}
}
func (h *HandlerOrder) AddOrder(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("%+v\n", qp)
	validate := validator.New()
	var errValidate []string
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, "Internal Server Error")
		json.NewEncoder(w).Encode(response)
		return
	}
	err := validate.VarWithValue(order.Email, order.EmailConfirm, "eqfield")
	if err != nil {
		errValidate = append(errValidate, "Email tidak Sama")
	}
	err = validate.Struct(order)
	if err != nil {
		errValidate = append(errValidate, err.Error())
	}
	if len(errValidate) > 0 {
		response := utils.SetResponse(w, model.Response{Error: errValidate}, http.StatusBadRequest, "Bad Request")
		json.NewEncoder(w).Encode(response)
		return
	}
	err = h.Service.AddOrder(&order)
	if err != nil {
		response := utils.SetResponse(w, model.Response{}, http.StatusInternalServerError, err.Error())
		json.NewEncoder(w).Encode(response)
		return
	}
	response := utils.SetResponse(w, model.Response{Data: order}, http.StatusCreated, "Order Berhasil Dibuat")
	json.NewEncoder(w).Encode(response)
}
