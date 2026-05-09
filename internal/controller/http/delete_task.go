package http

import (
	"encoding/json"
	"net/http"

	"proj/doollit/internal/dto"
	"proj/doollit/pkg/render"
)

func (h *Handlers) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// const op = "http.DeleteTask"

	var input dto.DeleteTaskInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.taskService.DeleteTask(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
