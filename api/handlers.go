package api

import (
	"io"
	"net/http"

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

	db.PostRequest(req)
}
