package api

import (
	"encoding/json"
	"net/http"

	"github.com/tousart/topz/internal/models"
)

type ProcResponse struct {
	Procs []models.Proc `json:"procs"`
}

func (ta *TopzApi) getProcHandler(w http.ResponseWriter, r *http.Request) {
	procs, err := ta.service.GetProc()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := ProcResponse{
		Procs: procs,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

// func getCPUHandler(w http.ResponseWriter, r *http.Request) {

// }

func (ta *TopzApi) WithHandlers(mux *Mux) {
	mux.HandleFunc(models.MGET, "/proc", ta.getProcHandler)
	// mx.HandleFunc(models.MGET, "/proc/cpu", getCPUHandler)
}
