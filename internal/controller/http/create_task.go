package http

import (
	"encoding/json"
	"net/http"

	"proj/doollit/internal/dto"
)

func (h *Handlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	input := dto.CreateTaskInput{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.taskService.CreateTask(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
}
