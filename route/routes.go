package route

import (
	"net/http"

	"github.com/codelotus/rivermq/handler"
)

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
		handler.CreateSubscriptionHandler,
	},
	Route{
		"GetSubscriptionByIDHandler",
		"GET",
		"/subscriptions/{subID}",
		handler.GetSubscriptionByIDHandler,
	},
	Route{
		"GetAllSubscriptionsHandler",
		"GET",
		"/subscriptions",
		handler.GetAllSubscriptionsHandler,
	},
	Route{
		"DeleteSubscriptionByIDHandler",
		"DELETE",
		"/subscriptions/{subID}",
		handler.DeleteSubscriptionByIDHandler,
	},
}
