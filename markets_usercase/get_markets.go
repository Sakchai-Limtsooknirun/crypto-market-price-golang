package markets_usercase

import (
	"alert/crypto/api"
	"alert/crypto/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetMarkets(w http.ResponseWriter, r *http.Request) {
	log.Print("[request] "+r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	var res *model.MarketCurrencyData = nil
	var ids = ""


	vars := mux.Vars(r)
	idsParam := vars["ids"]
	if idsParam != "" || len(idsParam) == 0 {
		ids = idsParam
	}
	query := r.URL.Query()
	idsQuery := query.Get("ids")
	if idsQuery != "" {
		ids = idsQuery
	}

	res = api.MarketPriceByIds(ids)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}
