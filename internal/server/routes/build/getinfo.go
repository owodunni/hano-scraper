package build

import (
	"encoding/json"
	"net/http"

	"github.com/owodunni/hano-scraper/internal/server/contenttype"
	"github.com/owodunni/hano-scraper/internal/server/httperr"
)

// Handler to get the program build information (GET /).
func (h *handler) getBuild(w http.ResponseWriter, r *http.Request) {
	_, responseContentType, err := contenttype.APICheck(r.Header)
	w.Header().Set("Content-Type", responseContentType)
	errResponder := httperr.NewResponder(responseContentType, h.logger)

	if err != nil {
		errResponder.Respond(w, http.StatusNotAcceptable, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(h.build)
	if err != nil {
		h.logger.Error(err.Error())
		errResponder.Respond(w, http.StatusInternalServerError, "")
		return
	}
}
