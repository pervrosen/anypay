/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: per.von.rosen@swedbank.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"testing"

	"github.com/gin-gonic/gin"
	openapi "github.com/pvr1/anypay/go"
)

func TestGetTransaction(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			openapi.GetTransaction(tt.args.c)
		})
	}
}

func TestGetTransactions(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			openapi.GetTransactions(tt.args.c)
		})
	}
}