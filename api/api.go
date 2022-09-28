package api

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	requestTimeout = 15 * time.Second
)

type (
	API struct {
		router http.Handler
	}

	profile struct {
		UserID string
	}

	apiHandle func(http.ResponseWriter, *http.Request, httprouter.Params, profile)
)

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	r = r.WithContext(ctx)

	api.router.ServeHTTP(w, r)
}

func NewAPI() (*API, error) {
	api := &API{}

	router := httprouter.New()

	router.GET("/requests", api.doAuthentication(api.GetAllRequests))
	router.GET("/profile/all", api.doAuthentication(api.getProfiles))
	router.GET("/profile", api.doAuthentication(api.getProfile))
	router.POST("/profile", api.doAuthentication(api.registerProfile))
	router.GET("/posts/:keyword", api.doAuthentication(api.mockHandler))
	router.GET("/grouppage/:title", api.doAuthentication(api.mockHandler))
	router.GET("/pagesearch/:keyword", api.doAuthentication(api.mockHandler))
	router.GET("/grouppage/:title/:postid", api.doAuthentication(api.mockHandler))
	router.POST("/grouppage/:title", api.doAuthentication(api.mockHandler))
	router.POST("/grouppage/:title/:postid", api.doAuthentication(api.mockHandler))
	router.POST("/grouppage/:title/:postid/:commentid", api.doAuthentication(api.mockHandler))

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, panicValue interface{}) {
		fmt.Printf("PANIC STACK TRACE\n%s\n", debug.Stack())
		err := fmt.Errorf("panicked with: %v", panicValue)
		respondWithError(w, r, err)
	}

	api.router = router

	return api, nil
}

func respondWithError(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
