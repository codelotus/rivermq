package main

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
	if err := model.SaveSubscription(sub); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(sub); err != nil {
		panic(err)
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
