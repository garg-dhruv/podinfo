package api

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/gorilla/mux"
)

// Log godoc
// @Summary Log text
// @Description logs the text in STDOUT
// @Tags HTTP API
// @Accept json
// @Produce json
// @Router /logtext/{text} [get]
// @Success 200 {object} api.MapResponse
func (s *Server) logtextHandler(w http.ResponseWriter, r *http.Request) {
	_, span := s.tracer.Start(r.Context(), "logtextHandler")
	defer span.End()

	vars := mux.Vars(r)

	text, ok := vars["text"]
	if !ok {
		s.ErrorResponse(w, r, span, "no text to log", http.StatusBadRequest)
		return
	}

	s.logger.Info("logging client text",
		zap.String("client-text", text))

	s.JSONResponse(w, r, map[string]string{"text": text})
}
