package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codelotus/rivermq/model"
	"github.com/gorilla/mux"
)

// CreateSubscriptionHandler does that
func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var sub model.Subscription
	json.NewDecoder(r.Body).Decode(&sub)
	resultSub, err := model.SaveSubscription(sub)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprint(w, err)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if err := json.NewEncoder(w).Encode(resultSub); err != nil {
			panic(err)
		}
	}
}

// GetAllSubscriptionsHandler does that
func GetAllSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetAllSubscriptions")
}

// GetSubscriptionByIDHandler does that
func GetSubscriptionByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subID := vars["subID"]
	fmt.Fprintf(w, "GetSubscription: %s", subID)
}

// DeleteSubscriptionByIDHandler does that
func DeleteSubscriptionByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	subID := vars["subID"]
	fmt.Fprintf(w, "DeleteSubscriptionByID: %s", subID)
}
