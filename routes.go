package main

import "net/http"

// Route defines the Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a slice of Routes
type Routes []Route

var routes = Routes{
	Route{
		"CreateSubscriptionHandler",
		"POST",
		"/subscriptions",
		CreateSubscriptionHandler,
	},
	Route{
		"GetSubscriptionByIDHandler",
		"GET",
		"/subscriptions/{subID}",
		GetSubscriptionByIDHandler,
	},
	Route{
		"GetAllSubscriptionsHandler",
		"GET",
		"/subscriptions",
		GetAllSubscriptionsHandler,
	},
	Route{
		"DeleteSubscriptionByIDHandler",
		"DELETE",
		"/subscriptions/{subID}",
		DeleteSubscriptionByIDHandler,
	},
}
