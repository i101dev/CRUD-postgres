package handlers

import (
	"net/http"

	"github.com/i101dev/rss-aggregator/util"
)

func HandleTest(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, 200, struct{}{})
}

func HandleError(w http.ResponseWriter, r *http.Request) {
	util.RespondWithError(w, 400, "Something went wrong")
}
