package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"proj/doollit/internal/dto"
	"proj/doollit/pkg/render"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("incorrect id: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)

		return
	}

	input := dto.UpdateTaskInput{ID: id}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	output, err := h.taskService.UpdateTask(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusNoContent)
}
