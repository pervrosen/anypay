/*
 * AnyPay
 *
 * This the AnyPay service targeted towards users
 *
 * API version: 1.0.0
 * Contact: pr.von.rosen@swedbank.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"
    "anypay/swagger"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

func main() {
	log.Printf("Server started")

	router := swagger.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
