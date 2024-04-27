package build

import (
	"net/http"

	"github.com/owodunni/hano-scraper/internal/server/middlewares/cors"
)

func (h *handler) options(w http.ResponseWriter, r *http.Request) {
	cors.AllowCORSMethods(r, w, http.MethodGet)
}
