package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MattGunson/my-server/db"
	"github.com/julienschmidt/httprouter"
)

func (api *API) mockHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user profile) {
	dat, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(dat)
	}

	headersMap := make(map[string]string)
	for h, vals := range r.Header {
		for _, val := range vals {
			w.Header().Add(h, val)
			headersMap[h] = val
		}
	}
	//w.WriteHeader(http.StatusOK)
	w.Write(dat)

	req := db.Request{
		Url:     r.URL.Path,
		Headers: headersMap,
		Body:    string(dat),
	}

	db.PostRequest(context.Background(), req)
}

func (api *API) registerProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user profile) {
	profile := db.Profile{
		Email:    r.Header["Email"][0],
		Name:     r.Header["Name"][0],
		Password: r.Header["Password"][0],
	}

	db.PostProfile(context.Background(), profile)
}

func (api *API) getProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user profile) {
	profile, err := db.GetProfile(context.Background(), r.Header["Email"][0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors with database connection: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	dat, err := json.Marshal(profile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors marshalling json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.Write(dat)
}

func (api *API) getProfiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user profile) {
	profiles, err := db.GetAllProfiles(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors with database connection: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	dat, err := json.Marshal(profiles)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors marshalling json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.Write(dat)
}

func (api *API) GetAllRequests(w http.ResponseWriter, r *http.Request, ps httprouter.Params, user profile) {
	profiles, err := db.GetAllRequests(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors with database connection: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	dat, err := json.Marshal(profiles)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Errors marshalling json: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	w.Write(dat)
}
