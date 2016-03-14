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
		"FindSubscriptionByIDHandler",
		"GET",
		"/subscriptions/{subID}",
		handler.FindSubscriptionByIDHandler,
	},
	Route{
		"FindAllSubscriptionsHandler",
		"GET",
		"/subscriptions",
		handler.FindAllSubscriptionsHandler,
	},
	Route{
		"DeleteSubscriptionByIDHandler",
		"DELETE",
		"/subscriptions/{subID}",
		handler.DeleteSubscriptionByIDHandler,
	},
}
