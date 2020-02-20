/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: pr.von.rosen@swedbank.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Openapi",
		"GET",
		"/.ambassador-internal/openapi-docs",
		Openapi,
	},

	Route{
		"Index",
		"GET",
		"/anypay/v1/",
		Index,
	},

	Route{
		"AddAccount",
		strings.ToUpper("Post"),
		"/anypay/v1/accounts",
		AddAccount,
	},

	Route{
		"GetAccount",
		strings.ToUpper("Get"),
		"/anypay/v1/accounts/{accountId}",
		GetAccount,
	},

	Route{
		"GetAccounts",
		strings.ToUpper("Get"),
		"/anypay/v1/accounts",
		GetAccounts,
	},

	Route{
		"UpdateAccount",
		strings.ToUpper("Put"),
		"/anypay/v1/accounts",
		UpdateAccount,
	},

	Route{
		"AddFxorder",
		strings.ToUpper("Post"),
		"/anypay/v1/fxorder",
		AddFxorder,
	},

	Route{
		"DeleteFxorder",
		strings.ToUpper("Delete"),
		"/anypay/v1/fxorder/{fxorderId}",
		DeleteFxorder,
	},

	Route{
		"GetFxorder",
		strings.ToUpper("Get"),
		"/anypay/v1/fxorder",
		GetFxorder,
	},

	Route{
		"GetFxorder2",
		strings.ToUpper("Get"),
		"/anypay/v1/fxorder/{fxorderId}",
		GetFxorder2,
	},

	Route{
		"UpdateFxorder",
		strings.ToUpper("Put"),
		"/anypay/v1/fxorder",
		UpdateFxorder,
	},

	Route{
		"AddPayment",
		strings.ToUpper("Post"),
		"/anypay/v1/payments",
		AddPayment,
	},

	Route{
		"DeletePayment",
		strings.ToUpper("Delete"),
		"/anypay/v1/payments/{paymentId}",
		DeletePayment,
	},

	Route{
		"GetPayment",
		strings.ToUpper("Get"),
		"/anypay/v1/payments/{paymentId}",
		GetPayment,
	},

	Route{
		"GetPayments",
		strings.ToUpper("Get"),
		"/anypay/v1/payments",
		GetPayments,
	},

	Route{
		"UpdatePayment",
		strings.ToUpper("Put"),
		"/anypay/v1/payments",
		UpdatePayment,
	},

	Route{
		"GetTransactins",
		strings.ToUpper("Get"),
		"/anypay/v1/transactions",
		GetTransactins,
	},

	Route{
		"AddUser",
		strings.ToUpper("Post"),
		"/anypay/v1/user",
		AddUser,
	},

	Route{
		"GetUser",
		strings.ToUpper("Get"),
		"/anypay/v1/users/{userId}",
		GetUser,
	},

	Route{
		"GetUses",
		strings.ToUpper("Get"),
		"/anypay/v1/user",
		GetUses,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/anypay/v1/user",
		UpdateUser,
	},
}
