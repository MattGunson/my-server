package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *API) doAuthentication(handler apiHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authnUser, err := api.authenticate(w, r)
		if err != nil {
			return
		}

		handler(w, r, ps, authnUser)
	}
}

func (api *API) authenticate(w http.ResponseWriter, r *http.Request) (profile, error) {
	return profile{UserID: "FakeUser"}, nil
}
