package http

import (
	"encoding/json"
	"net/http"

	"proj/doollit/internal/dto"
	"proj/doollit/pkg/render"
)

func (h *Handlers) ReadTask(w http.ResponseWriter, r *http.Request) {
	const op = "http.ReadTask"

	var input dto.ReadTaskInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.taskService.ReadTask(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
