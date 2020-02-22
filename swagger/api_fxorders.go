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
	"net/http"
)

func AddFxorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "FX Order place on market place\n")
}

func DeleteFxorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "FX Order deleted. You were so close...\n")
}

func GetFxorders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get FX Order list. Offer:100MM EUR/SEK @ 10.53 , Bid:100MM EUR/SEK @ 10.51 . \n")
}

func GetFxorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Get specific FX Order status. You are best offer: 100MM EUR/SEK @ 10.53 . \n")
}

func UpdateFxorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
