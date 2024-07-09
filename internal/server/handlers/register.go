package handlers

import (
	"io"
	"net/http"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/auth"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/serializers"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/store"
)

var Register http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, messageInternalServerError, http.StatusInternalServerError)
		return
	}
	r, err := serializers.NewRegistrationFromJSON(data)
	if err != nil {
		http.Error(res, "неверный формат запроса", http.StatusBadRequest)
		return
	}
	user, err := store.RegisterUser(r)
	if err != nil {
		http.Error(res, "логин уже занят", http.StatusConflict)
		return
	}
	err = auth.SetAuthorization(&res, user)
	if err != nil {
		http.Error(res, messageInternalServerError, http.StatusInternalServerError)
		return
	}
}
