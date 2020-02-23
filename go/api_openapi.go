/*
 * AnyPay
 *
 * This the AnyPay service targeted towards, parents with children doing payments and russian oligarks
 *
 * API version: 1.0.0
 * Contact: per.von.rosen@swedbank.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// OpenAPIjson - Get OpenAPI 3.0 JSON
func OpenAPIjson(c *gin.Context) {
	file, err := os.Open("/openapi.json")
	defer file.Close()
	if err != nil {
		c.String(http.StatusOK, "Could not find openapi file")
		return
	}

	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(http.StatusOK, "Could not read openapi file after loading it")
		return
	}

	c.String(http.StatusOK, string(filecontent))
}

// OpenAPIyaml - Get OpenAPI 3.0 YAML
func OpenAPIyaml(c *gin.Context) {
	file, err := os.Open("/openapi.yaml")
	defer file.Close()
	if err != nil {
		c.String(http.StatusOK, "Could not find openapi file")
		return
	}

	filecontent, err := ioutil.ReadAll(file)
	if err != nil {
		c.String(http.StatusOK, "Could not read openapi file after loading it")
		return
	}
	c.String(http.StatusOK, string(filecontent))

}