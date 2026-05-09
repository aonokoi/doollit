package http

import (
	"net/http"
	"strconv"

	"proj/doollit/internal/dto"
	"proj/doollit/pkg/render"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) ReadTask(w http.ResponseWriter, r *http.Request) {
	// const op = "http.ReadTask"

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	input := dto.ReadTaskInput{ID: id}

	output, err := h.taskService.ReadTask(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
