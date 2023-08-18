package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chebotarmv/gymshark/server"
	"github.com/chebotarmv/gymshark/services/calculator"
)

func CalculatePacks(w http.ResponseWriter, r *http.Request) {
	itemStr := r.URL.Query().Get("item")
	if itemStr == "" {
		http.Error(w, "Missing item value", http.StatusBadRequest)
		return
	}

	item, err := strconv.Atoi(itemStr)
	if err != nil {
		http.Error(w, "Invalid item value", http.StatusBadRequest)
		return
	}

	neededPacks := calculator.FindPerfectFit(server.PackSizes, item)
	respMap := make(map[int]int)
	for _, pack := range neededPacks {
		respMap[pack]++
	}

	err = json.NewEncoder(w).Encode(respMap)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
