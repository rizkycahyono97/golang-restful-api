package middleware

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Cek API KEY
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		//OK
		middleware.Handler.ServeHTTP(w, r)
	} else {
		//error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
