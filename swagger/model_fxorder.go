/*Package swagger AnyPay
 *
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
	"time"
)

/*Fxorder - a struct holding the FX order data necessary to IDentify and create an FX order*/
type Fxorder struct {
	ID int64 `json:"ID,omitempty"`

	FX int64 `json:"FX,omitempty"`

	Account *Account `json:"account,omitempty"`

	Quantity int32 `json:"quantity,omitempty"`

	OrderDate time.Time `json:"orderDate,omitempty"`

	SettlementDate time.Time `json:"settlementDate,omitempty"`
	// Order Status
	Status string `json:"status,omitempty"`

	Complete bool `json:"complete,omitempty"`
}
